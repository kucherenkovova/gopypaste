package xerrors

import (
	"errors"
	"slices"
)

// IsTimeout reports whether the provided error indicates a timeout condition.
func IsTimeout(err error) bool {
	if err == nil {
		return false
	}

	if te, ok := err.(interface{ Timeout() bool }); ok && te.Timeout() {
		return true
	}

	if joined, ok := err.(interface{ Unwrap() []error }); ok {
		return slices.ContainsFunc(joined.Unwrap(), IsTimeout)
	}

	inner := errors.Unwrap(err)

	return inner != nil && IsTimeout(inner)
}
