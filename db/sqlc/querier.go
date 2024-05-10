// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateUserWatch(ctx context.Context, arg CreateUserWatchParams) (UserWatch, error)
	DeleteUserFromWatch(ctx context.Context, userID uuid.UUID) error
	GetUserWatch(ctx context.Context, userID uuid.UUID) (UserWatch, error)
	GetUserWatchForUpdate(ctx context.Context, userID uuid.UUID) (UserWatch, error)
	ListUsersOnWatch(ctx context.Context, arg ListUsersOnWatchParams) ([]UserWatch, error)
	UpdateUserWatch(ctx context.Context, arg UpdateUserWatchParams) (UserWatch, error)
}

var _ Querier = (*Queries)(nil)
