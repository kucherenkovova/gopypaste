package xhttp

import (
	"context"
	"testing"
	"time"
)

func TestListenAndServeContextSuccess(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Millisecond)

	err := ListenAndServeContext(ctx, ":8080", nil)
	if err != nil {
		t.Errorf("err should be nil, got %s", err)
	}
}
