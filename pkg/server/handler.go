package server

import (
	"context"

	"github.com/iwanhae/kaustar/api"
)

var _ api.Handler = (*Server)(nil)

type Server struct{}

// CreateUser implements api.Handler
func (*Server) CreateUser(ctx context.Context, req *api.UserCreateRequest) (*api.User, error) {
	panic("unimplemented")
}

// ListUsers implements api.Handler
func (*Server) ListUsers(ctx context.Context, params api.ListUsersParams) (*api.UserList, error) {
	panic("unimplemented")
}

// NewError implements api.Handler
func (*Server) NewError(ctx context.Context, err error) *api.MessageStatusCode {
	panic("unimplemented")
}

// Whoami implements api.Handler
func (*Server) Whoami(ctx context.Context) (*api.User, error) {
	panic("unimplemented")
}
