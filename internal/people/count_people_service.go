package people

import (
	"context"
)

func (service *PeopleService) Count(ctx context.Context) (int, error) {
	count, err := service.Repository.Count(ctx)
	if err != nil {
		return count, err
	}
	return count, nil

}
