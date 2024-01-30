package people

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type PeopleServiceInterface interface {
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, p *People) (string, error)
	FindById(ctx context.Context, id string) (*People, error)
	Find(ctx context.Context, query string) (*[]People, error)
}

type PeopleService struct {
	Repository PeopleRepository
	Cache      *redis.Client
	Channel    chan<- *People
}

func NewService(pr PeopleRepository, cache *redis.Client, ch chan<- *People) *PeopleService {
	return &PeopleService{
		Repository: pr,
		Cache:      cache,
		Channel:    ch,
	}
}
