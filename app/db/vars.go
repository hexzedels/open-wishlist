package db

const (
	queryUser  = "SELECT * FROM users WHERE id=$1"
	insertUser = "INSERT INTO users (id, username, state) values ($1, $2, $3)"
)

const (
	queryInitUsers = `CREATE TABLE IF NOT EXISTS users (
		id bigint NOT NULL,
		username text NOT NULL,
		state text NOT NULL, 
		PRIMARY KEY (id)
	  )`
)
