package people

import (
	"context"
)

func (client *PeoplePostgres) Count(ctx context.Context) (int, error) {
	var count int
	err := client.DB.QueryRow("select count(*) from public.people").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
