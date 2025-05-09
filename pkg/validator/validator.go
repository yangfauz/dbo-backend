package validator

import (
	"bytes"
	"dbo-backend/pkg/response"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en_US"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

type ApiError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateRequest(req *http.Request, referenceStruct interface{}) (res response.Response, err error) {

	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		return response.Error(response.StatusUnprocessableEntity, ERR_INPUT, []string{ERR_EMPTY}), err
	}

	req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if len(bodyBytes) == 0 {
		return response.Error(response.StatusUnprocessableEntity, ERR_INPUT, []string{ERR_EMPTY}), errors.New(ERR_EMPTY)
	}

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(referenceStruct); err != nil {
		var errValue interface{}
		text := ERR_DATATYPE
		switch err.(type) {
		case *json.UnmarshalTypeError:
			typeError := err.(*json.UnmarshalTypeError)
			text = ERR_INPUT
			errValue = []ApiError{{typeError.Field, fmt.Sprintf("expected type %s, got %s", typeError.Type, typeError.Value)}}
		default:
			text = ERR_WRONG
			errValue = []string{fmt.Sprintf("failed when parsing body: %v", err)}
		}
		return response.Error(response.StatusUnprocessableEntity, text, errValue), errors.New(ERR_DATATYPE)
	}

	val := reflect.ValueOf(referenceStruct)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return response.Error(response.StatusUnprocessableEntity, ERR_WRONG, []string{"referenceStruct should be a non-nil pointer"}), errors.New(ERR_WRONG)
	}

	ptrType := reflect.TypeOf(referenceStruct)
	if ptrType.Elem().Kind() != reflect.Slice {
		res := validateStruct(referenceStruct)
		if res != nil {
			return response.Error(response.StatusUnprocessableEntity, ERR_INPUT, res), errors.New(ERR_VALIDATE)
		}
		return response.Error(response.StatusUnprocessableEntity, ACC_INPUT, res), nil
	}

	for i := 0; i < val.Elem().Len(); i++ {
		element := val.Elem().Index(i)
		if err := validateStruct(element.Interface()); err != nil {
			return response.Error(response.StatusUnprocessableEntity, ACC_INPUT, err), nil
		}
	}

	return response.Error(response.StatusUnprocessableEntity, ACC_INPUT, err), nil
}

func validateStruct(referenceStruct interface{}) (res interface{}) {

	validate := validator.New()

	//translation
	en := en_US.New()
	uni := ut.New(en, en, en)
	transEn, _ := uni.GetTranslator("en")
	enTranslations.RegisterDefaultTranslations(validate, transEn)

	useJsonFieldValidation(validate)

	// Validate the fields of the structure using validator/v10
	if err := validate.Struct(referenceStruct); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {

			errors := err.(validator.ValidationErrors)
			out := make([]interface{}, len(errors))
			for i, e := range errors {
				out[i] = ApiError{e.Field(), msgForTag(e.Tag(), e.Param(), e.Translate(transEn))}
			}

			return out
		}
	}

	return res
}

func msgForTag(tag string, param string, msgError string) string {
	switch tag {
	case "required_if":
		return "This field is required"
	}
	return msgError
}

func useJsonFieldValidation(v *validator.Validate) {
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}
