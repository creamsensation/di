package di

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	t.Run(
		"singleton", func(t *testing.T) {
			assert.Equal(t, configSingleton, Singleton().(config).configType)
		},
	)
}
