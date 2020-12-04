package apperrors

import (
	"errors"
	"net/http"
)

type ErrorModel struct {
	Message string
	Code    int
}

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrDatabaseConnection  = errors.New("Unable to connect to the Database")
	ErrDataNotFound        = errors.New("Failed to find requested information")
	ErrActionFailed        = errors.New("Failed to complete requested action")
	ErrBadRequest          = errors.New("Failed to parse body")
	ErrDatabaseRecord      = errors.New("Information is distorted")
)

func assertError(err error) *ErrorModel {
	if errors.Is(err, ErrInternalServerError) {
		return &ErrorModel{
			Message: ErrInternalServerError.Error(),
			Code:    500,
		}
	}
	if errors.Is(err, ErrDatabaseRecord) {
		return &ErrorModel{
			Message: ErrDatabaseRecord.Error(),
			Code:    417,
		}
	}
	if errors.Is(err, ErrDatabaseConnection) {
		return &ErrorModel{
			Message: ErrDatabaseConnection.Error(),
			Code:    500,
		}
	}
	if errors.Is(err, ErrDataNotFound) {
		return &ErrorModel{
			Message: ErrDataNotFound.Error(),
			Code:    404,
		}
	}
	if errors.Is(err, ErrBadRequest) {
		return &ErrorModel{
			Message: ErrBadRequest.Error(),
			Code:    400,
		}
	}
	if errors.Is(err, ErrActionFailed) {
		return &ErrorModel{
			Message: ErrActionFailed.Error(),
			Code:    501,
		}
	}
	return &ErrorModel{
		Message: "Unidentified Error",
		Code:    500,
	}
}

func ErrorResponse(e error, w http.ResponseWriter) {
	err := assertError(e)
	http.Error(w, err.Message, err.Code)
	return
}
