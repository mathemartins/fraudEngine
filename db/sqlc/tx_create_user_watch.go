package db

import "context"

type CreateUserWatchTxParams struct {
	CreateUserWatchParams
	AfterCreate func(user UserWatch) error
}

type CreateUserWatchTxResult struct {
	UserWatch UserWatch
}

func (store *SQLStore) CreateUserWatchTx(ctx context.Context, arg CreateUserWatchTxParams) (CreateUserWatchTxResult, error) {
	var result CreateUserWatchTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.UserWatch, err = q.CreateUserWatch(ctx, arg.CreateUserWatchParams)
		if err != nil {
			return err
		}

		return arg.AfterCreate(result.UserWatch)
	})

	return result, err
}
