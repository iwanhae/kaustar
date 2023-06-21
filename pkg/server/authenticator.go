package server

import (
	"context"

	"github.com/iwanhae/kaustar/api"
)

var _ api.SecurityHandler = (*Authenticator)(nil)

type Authenticator struct{}

// HandleBasicAuth implements api.SecurityHandler
func (*Authenticator) HandleBasicAuth(ctx context.Context, operationName string, t api.BasicAuth) (context.Context, error) {
	panic("unimplemented")
}

// HandleCookieAuth implements api.SecurityHandler
func (*Authenticator) HandleCookieAuth(ctx context.Context, operationName string, t api.CookieAuth) (context.Context, error) {
	panic("unimplemented")
}
