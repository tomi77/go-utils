package p_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomi77/go-utils/p"
)

func TestPtr(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		i := 42
		result := p.Ptr(i)
		assert.Equal(t, &i, result)
	})

	t.Run("string", func(t *testing.T) {
		s := "hello"
		result := p.Ptr(s)
		assert.Equal(t, &s, result)
	})

	t.Run("slice of int", func(t *testing.T) {
		slice := []int{1, 2, 3}
		result := p.Ptr(slice)
		assert.Equal(t, &slice, result)
	})

	t.Run("map of string to int", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		result := p.Ptr(m)
		assert.Equal(t, &m, result)
	})
}
