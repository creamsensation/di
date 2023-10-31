package di

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestProvide(t *testing.T) {
	t.Run(
		"provide service from container", func(t *testing.T) {
			provider := &testService{}
			c := New(
				Register[testService](provider),
			)
			p := Provide[testService](c)
			assert.Equal(t, testValue, p.value)
		},
	)
	t.Run(
		"provide and call function", func(t *testing.T) {
			provider := &testPtrService{}
			c := New(
				Register[testPtrService](provider, Singleton()),
			)
			p := Provide[testPtrService](c)
			p.testMethod()
			assert.Equal(t, testValue, p.value)
		},
	)
	t.Run(
		"provide with name", func(t *testing.T) {
			provider1 := &testService{value: 1}
			provider2 := &testService{value: 2}
			c := New(
				Register[testService](provider1, Name("first"), Singleton()),
				Register[testService](provider2, Name("second"), Singleton()),
			)
			p1 := Provide[testService](c, Name("first"))
			p2 := Provide[testService](c, Name("second"))
			assert.Equal(t, 1, p1.value)
			assert.Equal(t, 2, p2.value)
		},
	)
	t.Run(
		"provide nested interface", func(t *testing.T) {
			c := New(
				Register[testService](&testService{}),
				Register[nestedService](&nestedService{}),
			)
			p := Provide[testService](c)
			assert.Equal(t, testNestedValue, p.nestedService.getValue())
		},
	)
}
