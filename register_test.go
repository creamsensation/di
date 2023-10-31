package di

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	t.Run(
		"dependency", func(t *testing.T) {
			dep := Register[testService](&testService{value: testValue}).(dependency)
			assert.Equal(t, testValue, dep.provider.(*testService).value)
		},
	)
}
