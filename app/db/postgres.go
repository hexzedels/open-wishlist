package db

import (
	"context"
	"openwishlist/app/sdk/models"

	"database/sql"

	_ "github.com/lib/pq"
)

type IClient interface {
	GetUser(ctx context.Context, user *models.User) error
	CreateUser(ctx context.Context, user *models.User) error
}

type PostgresClient struct {
	conn *sql.Conn
}

func NewPostgresClient(ctx context.Context, connStr string) *PostgresClient {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	conn, err := db.Conn(ctx)
	if err != nil {
		panic(err)
	}

	_, err = conn.ExecContext(ctx, queryInitUsers)
	if err != nil {
		panic(err)
	}

	return &PostgresClient{
		conn: conn,
	}
}

func (r *PostgresClient) GetUser(ctx context.Context, user *models.User) error {
	row := r.conn.QueryRowContext(ctx, queryUser, user.ID)

	if err := row.Scan(user); err != nil {
		return err
	}

	return nil
}

func (r *PostgresClient) CreateUser(ctx context.Context, user *models.User) error {
	_, err := r.conn.ExecContext(ctx, insertUser, user.ID, user.Username, user.State)
	if err != nil {
		return err
	}

	return nil
}
