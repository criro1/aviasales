package httperrors

import (
	"fmt"
)

const (
	failedToDecodeJSON                     = "failed to decode JSON response: %s"
	failedToEncodeJSON                     = "failed to encode JSON response: %s"
	failedToEncodeQueryTypeIntErrorCreator = "failed to encode query type int response: %s"
)

// DecodeJSONErrorCreator ...
func DecodeJSONErrorCreator(err error) error {
	return fmt.Errorf(failedToDecodeJSON, err)
}

// EncodeJSONErrorCreator ...
func EncodeJSONErrorCreator(err error) error {
	return fmt.Errorf(failedToEncodeJSON, err)
}

// EncodeQueryTypeIntErrorCreator ...
func EncodeQueryTypeIntErrorCreator(err error) error {
	return fmt.Errorf(failedToEncodeQueryTypeIntErrorCreator, err)
}

// NewError ...
func NewError(status int, format string, v ...interface{}) error {
	return &httpError{
		Code:    status,
		Message: fmt.Sprintf(format, v...),
	}
}
