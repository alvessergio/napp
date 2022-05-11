package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

// the sentinel errors as real constants and immutable.
type err string

// NewError is just a syntax suggar for fmt.Error("%w: %s", err, cause)
// Used to give more context to the err error.
func NewError(err error, cause string) error {
	return fmt.Errorf("%w: %s", err, cause)
}

func (e err) Error() string {
	return string(e)
}

// We declare this custom error type so we can create
const (
	// ErrEmptyParams should be used when one or more required params are empty or nil
	ErrEmptyParams = err("one or more required parameters are empty or nil")
	// ErrItemNotFound is used when an Item is not found
	ErrItemNotFound = err("item not found")
)

func IsEqualError(typeErr, returnedErr error) bool {
	splitedError := strings.Split(returnedErr.Error(), ":")
	if len(splitedError) == 0 || splitedError[0] != typeErr.Error() {
		return false
	}

	return true
}

type errResponse struct {
	Err string `json:"error"`
}

func handleError(err error) (statusCode int, msg string) {
	switch {
	case IsEqualError(ErrEmptyParams, err):
		return http.StatusBadRequest, err.Error()
	case IsEqualError(ErrItemNotFound, err):
		return http.StatusBadRequest, "product not found"
	default:
		return http.StatusInternalServerError, "server.internal_error"
	}
}

func encodeErrorResponse(rw http.ResponseWriter, traceID string, err error) {
	if err == nil {
		panic("encodeError with nil error")
	}

	statusCode, msg := handleError(err)
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	e := errResponse{Err: msg}
	errResp, jsonErr := json.Marshal(e)
	if jsonErr != nil {
		log.WithFields(log.Fields{
			"event":  "error_serialize_failed",
			"reason": err,
		}).Error("encode_error_response_failed")
		return
	}
	rw.WriteHeader(statusCode)
	rw.Write(errResp)
}
