package db

import (
	"context"
	"github.com/jmoiron/sqlx"
    _ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresqlDB struct {
	*sqlx.DB
}

type PostgresqlTx struct {
	*sqlx.Tx
	context.Context
}

func NewPostgresqlDB(dsn string) (*PostgresqlDB, error) {
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, err
	}

    
    if err := db.Ping(); err != nil {
        return nil, err
    }

	postgresqlDb := &PostgresqlDB{db}

	return postgresqlDb, nil
}

func (db *PostgresqlDB) StartTransaction(ctx context.Context) (Transaction, error) {
	tx, err := db.DB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &PostgresqlTx{
		Tx:      tx,
		Context: ctx,
	}, nil
}
