package di

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestUtil(t *testing.T) {
	t.Run(
		"call provide", func(t *testing.T) {
			provider := &testService{}
			c := New(
				Register[testService](provider),
			)
			provider = callProvide[testService](
				c.(*container),
				dependency{
					provider: provider,
				},
			)
			assert.Equal(t, testValue, provider.value)
		},
	)
	t.Run(
		"find dependency with type", func(t *testing.T) {
			c := New(
				Register[testService](&testService{value: testValue}, Singleton()),
			)
			dep := findDependencyWithType[testService](c.(*container))
			assert.Equal(t, testValue, dep.provider.(*testService).value)
		},
	)
	t.Run(
		"is singleton", func(t *testing.T) {
			provider := &testService{value: testValue}
			c := New(
				Register[testService](provider, Singleton()),
			)
			assert.Equal(t, true, c.(*container).deps[0].config.singleton)
		},
	)
}
