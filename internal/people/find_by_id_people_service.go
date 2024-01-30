package people

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

func (s *PeopleService) FindById(ctx context.Context, id string) (*People, error) {
	cacheKey := id
	cachedData, err := s.Cache.Get(ctx, cacheKey).Result()
	if err != redis.Nil && cachedData != "" {
		data := &People{}
		json.Unmarshal([]byte(cachedData), &data)
		return data, nil
	}
	result, err := s.Repository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
