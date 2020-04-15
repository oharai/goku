package password

import (
	"context"

	"github.com/garupanojisan/goku/authz"
	"github.com/garupanojisan/goku/models"
)

type Auth struct {
}

func (a *Auth) Check(ctx context.Context) bool {
	return false
}

func (a *Auth) User(ctx context.Context) models.User {
	return nil
}

func (a *Auth) Attempt(ctx context.Context, credentials []*authz.Credential) bool {
	return false
}

func (a *Auth) Logout(ctx context.Context) error {
	return nil
}
