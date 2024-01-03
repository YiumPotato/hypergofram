package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

var QueryMaker *Queries
var conn *pgx.Conn

func InitDB() error {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "user=root password=secret123 dbname=hypergofram host=postgres sslmode=disable")
	if err != nil {
		return err
	}

	QueryMaker = New(conn)

	return nil
}

func CloseDB() {
	conn.Close(context.Background())
}
