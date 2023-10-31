package example

import (
	"github.com/creamsensation/di"
)

type PageRepository interface {
	GetPagesCount() int
}

type pageRepository struct {
	mock *dbMock
}

func (r *pageRepository) Provide(c di.Container) *pageRepository {
	return &pageRepository{
		mock: di.Provide[dbMock](c),
	}
}

func (r *pageRepository) GetPagesCount() int {
	return r.mock.count
}
