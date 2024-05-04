// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package query

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type PasswordResetToken struct {
	ID        uuid.UUID
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
	IsUsed    bool
	Token     string
	ExpiresAt pgtype.Timestamp
	UserID    uuid.UUID
}

type SsoProviderUser struct {
	ID             uuid.UUID
	CreatedAt      pgtype.Timestamp
	UpdatedAt      pgtype.Timestamp
	ProviderName   string
	UserID         uuid.UUID
	ProviderUserID string
}

type User struct {
	ID           uuid.UUID
	CreatedAt    pgtype.Timestamp
	UpdatedAt    pgtype.Timestamp
	Email        string
	PasswordHash *string
}