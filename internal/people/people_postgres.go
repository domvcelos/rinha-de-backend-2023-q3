package people

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
)

type PeoplePostgres struct {
	DB    *sql.DB
	Cache *redis.Client
}

func NewPostgres(db *sql.DB, cache *redis.Client) *PeoplePostgres {
	return &PeoplePostgres{DB: db, Cache: cache}
}
