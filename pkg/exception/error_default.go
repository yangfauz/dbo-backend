package exception

import "dbo-backend/pkg/response"

func ErrorUnauthorizedMessage(message string) Error {
	return Error{
		Status:  response.StatusUnauthorized,
		Message: message,
		Errors:  ErrUnauthorized,
	}
}

func ErrorBadRequestMessage(message string) Error {
	return Error{
		Status:  response.StatusBadRequest,
		Message: message,
		Errors:  ErrBadRequest,
	}
}

func ErrorBadRequest() Error {
	return Error{
		Status:  response.StatusBadRequest,
		Message: "Something Wrong",
		Errors:  ErrBadRequest,
	}
}

func ErrorLoginUnauthorized() Error {
	return Error{
		Status:  response.StatusUnauthorized,
		Message: "Invalid Email / Password",
		Errors:  ErrUnauthorized,
	}
}

func ErrorTokenNotValid() Error {
	return Error{
		Status:  response.StatusUnauthorized,
		Message: "Token Not Valid",
		Errors:  ErrUnauthorized,
	}
}
