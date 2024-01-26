package people

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

func (client *PeoplePostgres) FindById(ctx context.Context, id string) (*People, error) {
	cacheKey := "findById-" + id
	cachedData, err := client.Cache.Get(ctx, cacheKey).Result()
	p := People{}
	if err == redis.Nil && cachedData != "" {
		json.Unmarshal([]byte(cachedData), &p)
		return &p, nil
	}
	stmt, err := client.DB.Prepare(`select p.apelido, p.id, p.nome, p.nascimento, p.stack from public.people p where p.id = $1;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var stacks string
	err = stmt.QueryRow(id).Scan(&p.Apelido, &p.Id, &p.Nome, &p.Nascimento, &stacks)
	if err != nil {
		return nil, err
	}
	if stacks != "" {
		p.Stack = strings.Split(stacks, " ")
	}
	dataToCache, err := json.Marshal(p)
	if err != nil {
		fmt.Print("Failed to marshal resp")
	}
	err = client.Cache.Set(ctx, cacheKey, dataToCache, 5*time.Minute).Err()
	if err != nil {
		fmt.Print("Failed to set cache")
	}
	return &p, nil
}
