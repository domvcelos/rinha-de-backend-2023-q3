package people

import (
	"context"
)

type PeopleRepository interface {
	Create(ctx context.Context, p *People) (string, error)
	FindById(ctx context.Context, id string) (*People, error)
	Count(ctx context.Context) (int, error)
	Find(ctx context.Context, query string) (*[]People, error)
}
