package helloworld

import (
	"context"

	"github.com/KadyrPoyraz/httplayout/internal/repository/db"
)

type repo struct {
	databse *db.PostgresqlDB
}

func NewPostgresqlRepo(database *db.PostgresqlDB) Repo {
	return &repo{
		databse: database,
	}
}

func (r *repo) GetHelloWorld(ctx context.Context) ([]string, error) {
	query := `
        SELECT name FROM users
    `
	rows, err := r.databse.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}

    var text []string
	for rows.Next() {
		var t string
		err := rows.Scan(&t)
        if err != nil {
            return nil, err
        }
        text = append(text, t)
	}

	return text, nil
}
