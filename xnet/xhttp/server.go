package xhttp

import (
	"context"
	"errors"
	"net/http"
)

// ListenAndServeContext works just like ListenAndServe with 2 additions:
// - it respects the context and closes the server when the context is done
// - it doesn't return an annoying http.ErrServerClosed error when the server is closed
func ListenAndServeContext(ctx context.Context, addr string, handler http.Handler) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	server := &http.Server{Addr: addr, Handler: handler}

	go func() {
		<-ctx.Done()
		_ = server.Shutdown(context.Background())
	}()

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
