package xerrors

import (
	"errors"
)

// IsTimeout reports whether the provided error indicates a timeout condition.
func IsTimeout(err error) bool {
	if err == nil {
		return false
	}

	e, ok := err.(interface{ Timeout() bool })
	if ok && e.Timeout() {
		return true
	}

	inner := errors.Unwrap(err)

	return inner != nil && IsTimeout(inner)
}
