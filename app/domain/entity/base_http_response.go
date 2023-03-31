package entity

import (
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type BaseHttpResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type errorMessage struct {
	Messages []string `json:"messages"`
}

type BaseHttpResponseInterface interface {
	SuccessResponse(data interface{}) BaseHttpResponse
	ErrorFieldResponse(errors error, errorCode int) BaseHttpResponse
	ErrorMessageResponse(errors error, errorCode int) BaseHttpResponse
}

func (b BaseHttpResponse) SuccessResponse(data interface{}) BaseHttpResponse {
	return BaseHttpResponse{
		Success: true,
		Code:    200,
		Data:    data,
	}
}

func (b BaseHttpResponse) ErrorFieldResponse(err error, errorCode int) BaseHttpResponse {
	errorField := []string{}
	myMap := make(map[string][]string)

	for i, e := range err.(validator.ValidationErrors) {
		errF := strings.ToLower(e.Field())
		errorField = append(errorField, errF)
		errM := strings.ToLower(e.ActualTag())
		myMap[errorField[i]] = append(myMap[errorField[i]], errM)

	}

	return BaseHttpResponse{
		Success: false,
		Code:    errorCode,
		Errors:  myMap,
	}
}

func (b BaseHttpResponse) ErrorMessageResponse(err error, errorCode int) BaseHttpResponse {
	errorArr := []string{}
	errorArr = append(errorArr, http.StatusText(errorCode))

	return BaseHttpResponse{
		Success: false,
		Code:    errorCode,
		Errors:  errorMessage{Messages: errorArr},
	}
}
