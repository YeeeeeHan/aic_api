package helpers

import (
	"bytes"

	"github.com/oklog/ulid/v2"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func GenerateID() string {
	return ulid.Make().String()
}

func Ternary[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}

func CleanText(text string) string {
	return string(bytes.TrimSpace(bytes.Replace([]byte(text), newline, space, -1)))
}
