package errors

import (
	"fmt"

	"gitlab.com/aic/aic_api/biz/util/Status"
)

func Failed(status *Status.Status) bool {
	return !Succeeded(status)
}

func Succeeded(status *Status.Status) bool {
	return status == Status.Success
}

type InvalidParamsError struct {
	msg string
}

func (e *InvalidParamsError) Error() string {
	return "Invalid params: " + e.msg
}

func NewInvalidParamsError(format string, a ...interface{}) *InvalidParamsError {
	return &InvalidParamsError{fmt.Sprintf(format, a...)}
}

type InternalError struct {
	msg string
}

func NewInternalError(format string, a ...interface{}) *InternalError {
	return &InternalError{fmt.Sprintf(format, a...)}
}

func (e *InternalError) Error() string {
	return "Internal error: " + e.msg
}

type AuthorisationError struct {
	msg string
}

func NewAuthorisationError(format string, a ...interface{}) *AuthorisationError {
	return &AuthorisationError{fmt.Sprintf(format, a...)}
}

func (e *AuthorisationError) Error() string {
	return "Authorisation error: " + e.msg
}

func GetStatus(err error) *Status.Status {
	if err == nil {
		return Status.Success
	}
	switch err.(type) {
	case *InvalidParamsError:
		return Status.InvalidParams
	case *AuthorisationError:
		return Status.AuthorisationError
	case *EventPasscodeExistsError:
		return Status.EventPasscodeExistsError
	case *EventAlreadyJoinedError:
		return Status.EventAlreadyJoinedError
	case *EventAlreadyFullError:
		return Status.EventAlreadyFullError
	case *EventPasscodeInvalidError:
		return Status.EventPasscodeInvalidError
	default:
		return Status.Error
	}
}
