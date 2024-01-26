package people

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func (s *PeopleService) Create(ctx context.Context, p *People) (string, error) {
	cacheKey := strings.ToLower(p.Apelido)
	cachedData, err := s.Cache.Get(ctx, cacheKey).Result()
	if err != redis.Nil && cachedData != "" {
		return "", errors.New("nickname already exist")
	}
	id := (uuid.New()).String()
	p.Id = id
	s.Channel <- p
	err = s.Cache.Set(ctx, cacheKey, p.Apelido, 0).Err()
	if err != nil {
		fmt.Print("Failed to set cache")
	}
	return id, nil
}
