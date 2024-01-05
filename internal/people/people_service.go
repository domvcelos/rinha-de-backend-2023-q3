package people

import (
	"context"
)

type PeopleServiceInterface interface {
	Create(ctx context.Context, p *People) (string, error)
	FindById(ctx context.Context, id string) (*People, error)
	Count(ctx context.Context) (int, error)
	Find(ctx context.Context, query string) (*[]People, error)
}

type PeopleService struct {
	Repository PeopleRepository
}

func NewService(pr PeopleRepository) *PeopleService {
	return &PeopleService{
		Repository: pr,
	}
}

func (service *PeopleService) Create(ctx context.Context, p *People) (string, error) {
	id, err := service.Repository.Create(ctx, p)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (service *PeopleService) FindById(ctx context.Context, id string) (*People, error) {
	result, err := service.Repository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (service *PeopleService) Find(ctx context.Context, query string) (*[]People, error) {
	result, err := service.Repository.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (service *PeopleService) Count(ctx context.Context) (int, error) {
	count, err := service.Repository.Count(ctx)
	if err != nil {
		return count, err
	}
	return count, nil

}
