package db

const (
	queryUser  = "SELECT * FROM users WHERE id=$1"
	insertUser = "INSERT INTO users (id, username, state, wishlist_id) values ($1, $2, $3, $4)"
	updateUser = "UPDATE users SET state=$1, wishlist_id=$2 WHERE id=$3"
)

const (
	queryWishlist  = "SELECT * FROM wishlists WHERE id=$1"
	insertWishlist = "INSERT INTO wishlists (owner_id, name) values ($1, $2)"
	queryWishlists = "SELECT * FROM wishlists WHERE owner_id=$1"
)

const (
	queryInitUsers = `CREATE TABLE IF NOT EXISTS users (
		id bigint NOT NULL,
		username text NOT NULL,
		state text NOT NULL, 
		wishlist_id bigint NOT NULL,
		PRIMARY KEY (id)
	  )`
	queryInitWishlists = `
	CREATE TABLE IF NOT EXISTS wishlists (
		id bigserial NOT NULL,
		owner_id bigint NOT NULL,
		name text NOT NULL,
		PRIMARY KEY (id)
	  ); `
	queryInitItems = `
	CREATE TABLE IF NOT EXISTS items (
		id bigserial NOT NULL,
		wishlist_id bigint NOT NULL,
		name text NOT NULL,
		url text NOT NULL,
		PRIMARY KEY (id)
	  );`
)
