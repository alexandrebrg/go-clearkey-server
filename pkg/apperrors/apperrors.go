package apperrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type errorWrapper struct {
	Message  string `json:"message"`
	Code     string `json:"serviceCode"`
	HttpCode int    `json:"httpCode"`
	Data     string `json:"data"`
}

func (e *errorWrapper) String() string {
	return e.Message
}

func (e *errorWrapper) Marshal() string {
	rawMsg, err := json.Marshal(e)

	if err != nil {
		return fmt.Sprintf("Could not parse error %s", e.Code)
	}
	return string(rawMsg)
}

func newErrorWrapper(identifier string, errorText string, httpCode int) *errorWrapper {
	return &errorWrapper{
		Message:  errorText,
		Code:     identifier,
		HttpCode: httpCode,
	}
}

func (e *errorWrapper) WithData(data string) *errorWrapper {
	return &errorWrapper{
		Message:  e.Message,
		Code:     e.Code,
		HttpCode: e.HttpCode,
		Data:     data,
	}
}

func ReturnHttpError(w http.ResponseWriter, error *errorWrapper) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(error.HttpCode)
	fmt.Fprintln(w, error.Marshal())
}

var (
	/*
	 * HTTP CODES
	 */
	InvalidInput = newErrorWrapper("1", "input given is invalid", http.StatusBadRequest)
	NotFound = newErrorWrapper("2", "requested resource not found", http.StatusNotFound)
	Internal = newErrorWrapper("3", "an error occurred on our side", http.StatusInternalServerError)

	/*
	 * APPLICATION INTERNAL ERRORS
	 */
	EnvVarLoadFailed = errors.New("could not environment variables")
)