package exception

import (
	"database/sql"
	"dbo-backend/pkg/response"
)

func ErrorSqlNotFound(message string, err error) *Error {
	switch err {
	case nil:
		err = nil
	case sql.ErrNoRows:
		return &Error{
			Status:  response.StatusNotFound,
			Message: message,
			Errors:  ErrNotFound,
		}
	default:
		return &Error{
			Status:  response.StatusBadRequest,
			Message: "Something Wrong",
			Errors:  ErrBadRequest,
		}
	}
	return nil
}

func ErrorSqlCheckNotFound(err error) *Error {
	switch err {
	case nil:
		err = nil
	case sql.ErrNoRows:
		err = nil
	default:
		return &Error{
			Status:  response.StatusBadRequest,
			Message: "Something Wrong",
			Errors:  ErrBadRequest,
		}
	}
	return nil
}

func ErrorSqlConflict(message string, err error) *Error {
	switch err {
	case nil:
		return &Error{
			Status:  response.StatusConflicted,
			Message: message,
			Errors:  ErrConflicted,
		}
	case sql.ErrNoRows:
		err = nil
	default:
		return &Error{
			Status:  response.StatusBadRequest,
			Message: "Something Wrong",
			Errors:  ErrBadRequest,
		}
	}
	return nil
}
