package people

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type PeopleServiceMock struct {
	mock.Mock
}

func NewPeopleServiceMock() *PeopleServiceMock {
	return &PeopleServiceMock{
		mock.Mock{},
	}
}

func (m *PeopleServiceMock) Count(ctx context.Context) (int, error) {
	args := m.Called(ctx)
	return args.Int(0), args.Error(1)
}

func (m *PeopleServiceMock) Create(ctx context.Context, p *People) (string, error) {
	return "0", nil
}

func (m *PeopleServiceMock) FindById(ctx context.Context, id string) (*People, error) {
	return &People{}, nil
}
func (m *PeopleServiceMock) Find(ctx context.Context, query string) (*[]People, error) {
	return &[]People{}, nil
}
