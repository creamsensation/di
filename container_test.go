package di

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestContainer(t *testing.T) {
	t.Run(
		"create", func(t *testing.T) {
			c := New(
				Register[testService](&testService{value: testValue}, Singleton()),
			)
			assert.Equal(t, 1, len(c.(*container).deps))
		},
	)
}
