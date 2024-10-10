package errors

import "fmt"

type EventAlreadyFullError struct {
	msg string
}

func NewEventAlreadyFullError(format string, a ...interface{}) *EventAlreadyFullError {
	return &EventAlreadyFullError{fmt.Sprintf(format, a...)}
}

func (e *EventAlreadyFullError) Error() string {
	return "EventAlreadyFull error: " + e.msg
}

type EventPasscodeExistsError struct {
	msg string
}

func NewEventPasscodeExistsError(format string, a ...interface{}) *EventPasscodeExistsError {
	return &EventPasscodeExistsError{fmt.Sprintf(format, a...)}
}

func (e *EventPasscodeExistsError) Error() string {
	return "EventPasscodeExists error: " + e.msg
}

type EventAlreadyJoinedError struct {
	msg string
}

func NewEventAlreadyJoinedError(format string, a ...interface{}) *EventAlreadyJoinedError {
	return &EventAlreadyJoinedError{fmt.Sprintf(format, a...)}
}

func (e *EventAlreadyJoinedError) Error() string {
	return "EventAlreadyJoined error: " + e.msg
}

type EventPasscodeInvalidError struct {
	msg string
}

func NewEventPasscodeInvalidError(format string, a ...interface{}) *EventPasscodeInvalidError {
	return &EventPasscodeInvalidError{fmt.Sprintf(format, a...)}
}

func (e *EventPasscodeInvalidError) Error() string {
	return "EventPasscodeInvalid error: " + e.msg
}
