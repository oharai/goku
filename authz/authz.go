package authz

import (
	"context"

	"github.com/oharai/goku/models"
)

type Auth interface {
	// check if current user has logged in
	Check(ctx context.Context) bool
	// get current login user
	User(ctx context.Context) models.User
	// attempt authorization
	Attempt(ctx context.Context, credentials []*Credential) *Result
	// logout
	Logout(ctx context.Context) error
}

type Credential interface {
}

type Result interface {
}
