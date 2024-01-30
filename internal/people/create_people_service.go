package people

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func (s *PeopleService) Create(ctx context.Context, p *People) (string, error) {
	apleidoCacheKey := strings.ToLower(p.Apelido)
	cachedData, err := s.Cache.Get(ctx, apleidoCacheKey).Result()
	if err != redis.Nil && cachedData != "" {
		return "", errors.New("nickname already exist")
	}
	id := (uuid.New()).String()
	p.Id = id
	s.Channel <- p
	err = s.Cache.Set(ctx, apleidoCacheKey, p.Apelido, 0).Err()
	if err != nil {
		fmt.Print("Failed to set cache nickname")
	}
	dataToCache, err := json.Marshal(p)
	if err != nil {
		fmt.Print("Failed to marshal resp")
	}
	err = s.Cache.Set(ctx, id, dataToCache, 0).Err()
	if err != nil {
		fmt.Print("Failed to set cache people")
	}
	return id, nil
}
