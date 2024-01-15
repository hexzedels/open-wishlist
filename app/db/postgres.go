package db

type IClient interface {
}

type PostgresClient struct {
}

func NewPostgresClient() *PostgresClient {
	return &PostgresClient{}
}
