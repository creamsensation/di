package example

import "github.com/creamsensation/di"

type DashboardService interface {
	GetPagesMetrics() map[string]int
}

type dashboardService struct {
	pageRepository PageRepository
}

func (s *dashboardService) Provide(c di.Container) *dashboardService {
	return &dashboardService{
		pageRepository: di.Provide[pageRepository](c),
	}
}

func (s *dashboardService) GetPagesMetrics() map[string]int {
	return map[string]int{
		"count": s.pageRepository.GetPagesCount(),
	}
}
