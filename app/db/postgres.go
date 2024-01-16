package db

import (
	"context"
	"openwishlist/app/sdk/models"

	"database/sql"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type IClient interface {
	GetUser(ctx context.Context, user *models.User) error
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error

	CreateWishlist(ctx context.Context, wishlist *models.Wishlist) error
	ListWishlists(ctx context.Context, user *models.User) ([]*models.Wishlist, error)
}

type PostgresClient struct {
	conn   *sql.Conn
	logger *zap.Logger
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

	_, err = conn.ExecContext(ctx, queryInitWishlists)
	if err != nil {
		panic(err)
	}

	_, err = conn.ExecContext(ctx, queryInitItems)
	if err != nil {
		panic(err)
	}

	return &PostgresClient{
		conn: conn,
	}
}

func (r *PostgresClient) GetUser(ctx context.Context, user *models.User) error {
	row := r.conn.QueryRowContext(ctx, queryUser, user.ID)

	if err := row.Scan(&user.ID, &user.Username, &user.State); err != nil {
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

func (r *PostgresClient) UpdateUser(ctx context.Context, user *models.User) error {
	_, err := r.conn.ExecContext(ctx, updateUser, user.State, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresClient) CreateWishlist(ctx context.Context, wishlist *models.Wishlist) error {
	_, err := r.conn.ExecContext(ctx, insertWishlist, wishlist.ID, wishlist.OwnerID, wishlist.Name)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresClient) ListWishlists(ctx context.Context, user *models.User) ([]*models.Wishlist, error) {
	var out []*models.Wishlist

	rows, err := r.conn.QueryContext(ctx, queryWishlists, user.ID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		temp := &models.Wishlist{}

		if err := rows.Scan(&temp.ID, &temp.OwnerID, &temp.Name); err != nil {
			r.logger.Error("failed to scan row", zap.Error(err))
			continue
		}

		out = append(out, temp)
	}

	return out, nil
}
