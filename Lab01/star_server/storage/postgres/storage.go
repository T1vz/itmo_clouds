package postgres

import (
	"context"
	"fmt"
	"github.com/T1vz/itmo_clouds/Lab01/star_server/storage"
	"github.com/jackc/pgx/v5"
	"os"
)

type Repo struct {
	conn *pgx.Conn
}

func (r *Repo) DoRequest(request string) error {
	_, err := r.conn.Exec(context.Background(), request)

	return err
}

func New(url string) storage.IRepository {
	repo := &Repo{}

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	repo.conn = conn

	return repo
}
