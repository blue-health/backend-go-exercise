package account

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID           uuid.UUID `db:"id" json:"id"`
	Active       bool      `db:"active" json:"active"`
	Name         string    `db:"name" json:"name"`
	Email        string    `db:"email" json:"email"`
	PasswordHash []byte    `db:"password_hash" json:"-"`
	InsertedAt   time.Time `db:"inserted_at" json:"inserted_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}
