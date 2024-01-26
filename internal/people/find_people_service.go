package people

import (
	"context"
)

func (service *PeopleService) Find(ctx context.Context, query string) (*[]People, error) {
	result, err := service.Repository.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}
