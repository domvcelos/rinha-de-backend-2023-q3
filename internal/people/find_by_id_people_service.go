package people

import (
	"context"
)

func (service *PeopleService) FindById(ctx context.Context, id string) (*People, error) {
	result, err := service.Repository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
