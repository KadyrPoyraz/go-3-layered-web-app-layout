package db

import "context"

type Transaction interface {
	Commit() error
	Rollback() error
}

type DB interface {
	StartTransaction(context.Context) (Transaction, error)
}
