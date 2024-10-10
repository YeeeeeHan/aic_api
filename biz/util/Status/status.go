package Status

import (
	"fmt"
)

type Status struct {
	code int32
	msg  string
}

func (s *Status) Code() int32 {
	return s.code
}

func (s *Status) Msg() string {
	return s.msg
}

func (s *Status) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", s.code, s.Msg())
}

func (s *Status) Equal(t *Status) bool {
	if t == nil {
		return s == nil
	}
	if s == nil {
		return t == nil
	}
	return s.Code() == t.Code()
}
