package db

import "openwishlist/app/sdk/models"

type IClient interface {
	GetUser(user *models.User) error
}

type PostgresClient struct {
}

func NewPostgresClient() *PostgresClient {
	return &PostgresClient{}
}

func (r *PostgresClient) GetUser(user *models.User) error {
	return nil
}
