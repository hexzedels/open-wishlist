package db

const (
	queryUser  = "SELECT * FROM users WHERE id=$1"
	insertUser = "INSERT INTO users (id, username, state) values ($1, $2, $3)"
	updateUser = "UPDATE users SET state=$1 WHERE id=$2"
)

const (
	insertWishlist = "INSERT INTO wishlists (id, owner_id, name) values ($1, $2, $3)"
	queryWishlists = "SELECT * FROM wishlists WHERE owner_id=$1"
)

const (
	queryInitUsers = `CREATE TABLE IF NOT EXISTS users (
		id bigint NOT NULL,
		username text NOT NULL,
		state text NOT NULL, 
		PRIMARY KEY (id)
	  )`
	queryInitWishlists = `CREATE TABLE IF NOT EXISTS wishlists (
		id bigint NOT NULL,
		owner_id bigint NOT NULL,
		name text NOT NULL,
		PRIMARY KEY (id)
	  )`
	queryInitItems = `CREATE TABLE IF NOT EXISTS items (
		id bigint NOT NULL,
		wishlist_id bigint NOT NULL,
		name text NOT NULL,
		url text NOT NULL,
		PRIMARY KEY (id)
	  )`
)
