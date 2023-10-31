package example

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
	
	"github.com/creamsensation/di"
)

func TestExample(t *testing.T) {
	t.Run(
		"service use repository with mock data", func(t *testing.T) {
			value := 10
			c := di.New(
				di.Register[dbMock](di.CreateProvider(&dbMock{count: value}), di.Singleton()),
				di.Register[pageRepository](&pageRepository{}),
				di.Register[dashboardService](&dashboardService{}),
			)
			ds := di.Provide[dashboardService](c)
			assert.Equal(t, value, ds.GetPagesMetrics()["count"])
		},
	)
}
