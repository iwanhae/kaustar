// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CreateUser implements createUser operation.
//
// Create a new user.
//
// POST /users
func (UnimplementedHandler) CreateUser(ctx context.Context, req *UserCreateRequest) (r *User, _ error) {
	return r, ht.ErrNotImplemented
}

// ListUsers implements listUsers operation.
//
// List users.
//
// GET /users
func (UnimplementedHandler) ListUsers(ctx context.Context, params ListUsersParams) (r *UserList, _ error) {
	return r, ht.ErrNotImplemented
}

// Whoami implements whoami operation.
//
// Get the current user.
//
// GET /whoami
func (UnimplementedHandler) Whoami(ctx context.Context) (r *User, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *MessageStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *MessageStatusCode) {
	r = new(MessageStatusCode)
	return r
}
