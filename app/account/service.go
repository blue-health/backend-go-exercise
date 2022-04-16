package account

import (
	"context"

	"github.com/google/uuid"
)

type (
	Service interface {
		Login(context.Context, LoginCmd) (*Account, error)
		Get(context.Context, uuid.UUID) (*Account, error)
		GetActive(context.Context) ([]Account, error)
		UpdateName(context.Context, uuid.UUID, string) error
		Delete(context.Context, uuid.UUID) error
	}

	LoginCmd struct {
		Email    string
		Password string
	}
)
