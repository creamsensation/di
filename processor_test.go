package di

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestProcessor(t *testing.T) {
	t.Run(
		"config", func(t *testing.T) {
			cfg := processConfig(&testService{}, Name("test"), Singleton())
			assert.Equal(t, "test", cfg.name)
			assert.Equal(t, true, cfg.singleton)
		},
	)
}
