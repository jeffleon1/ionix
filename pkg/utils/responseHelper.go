package utils

import (
	"net/http"

	"github.com/jeffleon1/ionix/pkg/models"
)

//BaseResponseModel helpers

// custom response
func NewResponse(status int, data interface{}, message string, err interface{}) models.BaseResponseModel {
	return models.BaseResponseModel{
		Status:  status,
		Data:    data,
		Message: message,
		Error:   err,
	}
}

// returns http 200 OK
func StatusOK(data interface{}) models.BaseResponseModel {
	return models.BaseResponseModel{
		Status:  http.StatusOK,
		Data:    data,
		Message: "OK",
	}
}

// returns http 400
func StatusFail(err string) models.BaseResponseModel {
	return models.BaseResponseModel{
		Status:  http.StatusBadRequest,
		Message: "Something goes wrong",
		Error:   err,
	}
}

// returns http 422
func StatusUnprocessableEntity(err string) models.BaseResponseModel {
	return models.BaseResponseModel{
		Status:  http.StatusUnprocessableEntity,
		Message: "Something goes wrong",
		Error:   err,
	}
}

// retuns http 401
func StatusUnauthorized(err string) models.BaseResponseModel {
	return models.BaseResponseModel{
		Status:  http.StatusUnauthorized,
		Message: "Unauthorized",
		Error:   err,
	}
}

// returns http 500
func UnhandledError(err string) models.BaseResponseModel {
	return models.BaseResponseModel{
		Status:  http.StatusInternalServerError,
		Message: "Unhandled error occurred. Please try again later",
		Error:   err,
	}
}

// returns http 404
func StatusNotFound(err string) models.BaseResponseModel {
	return models.BaseResponseModel{
		Status:  http.StatusNotFound,
		Message: "Not found",
		Error:   err,
	}
}
