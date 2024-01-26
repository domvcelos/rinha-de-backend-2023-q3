package people

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

func (client *PeoplePostgres) Find(ctx context.Context, query string) (*[]People, error) {
	term := strings.ToLower(query)
	cacheKey := "find-" + term
	cachedData, err := client.Cache.Get(ctx, cacheKey).Result()
	peoples := []People{}
	if err == redis.Nil && cachedData != "" {
		json.Unmarshal([]byte(cachedData), &peoples)
		return &peoples, nil
	}
	stmt, err := client.DB.Prepare(`select p.apelido, p.id, p.nome, p.nascimento, p.stack from public.people p where p.busca_tgrm like '%'||$1||'%' LIMIT 50;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(term)

	for rows.Next() {
		p := People{}
		var stacks string
		err = rows.Scan(&p.Apelido, &p.Id, &p.Nome, &p.Nascimento, &stacks)
		if stacks != "" {
			p.Stack = strings.Split(stacks, " ")
		}
		peoples = append(peoples, p)
	}
	if err != nil {
		return nil, err
	}
	dataToCache, err := json.Marshal(peoples)
	if err != nil {
		fmt.Print("Failed to marshal resp")
	}
	err = client.Cache.Set(ctx, cacheKey, dataToCache, 10*time.Minute).Err()
	if err != nil {
		fmt.Print("Failed to set cache")
	}

	return &peoples, nil
}
