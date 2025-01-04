package xerrors_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/kucherenkovova/gopypaste/xerrors"
)

type customTimeoutError struct{}

func (e customTimeoutError) Error() string { return "timeout" }

func (e customTimeoutError) Timeout() bool { return true }

func TestIsTimeout(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{name: "nil", err: nil, want: false},
		{name: "context timeout", err: context.DeadlineExceeded, want: true},
		{name: "wrapped context timeout", err: fmt.Errorf("bad stuff: %w", context.DeadlineExceeded), want: true},
		{name: "custom timeout", err: customTimeoutError{}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xerrors.IsTimeout(tt.err); got != tt.want {
				t.Errorf("IsTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}
