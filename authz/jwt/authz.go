package jwt

import (
	"context"

	"github.com/oharai/goku/authz"
	"github.com/oharai/goku/models"
)

type Auth struct {
}

func (a *Auth) Check(ctx context.Context) bool {
	return false
}

func (a *Auth) User(ctx context.Context) models.User {
	return nil
}

func (a *Auth) Attempt(ctx context.Context, credentials []*authz.Credential) *authz.Result {
	return nil
}

func (a *Auth) Logout(ctx context.Context) error {
	return nil
}
