package s_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomi77/go-utils/s"
)

func TestMap(t *testing.T) {
	{
		tests := []struct {
			name     string
			input    []int
			fn       func(int) int
			expected []int
		}{
			{
				name:     "double each element",
				input:    []int{1, 2, 3},
				fn:       func(i int) int { return i * 2 },
				expected: []int{2, 4, 6},
			},
			{
				name:     "empty slice",
				input:    []int{},
				fn:       func(i int) int { return i * 2 },
				expected: []int{},
			},
			{
				name:     "nil function",
				input:    []int{1, 2, 3},
				fn:       nil,
				expected: nil,
			},
			{
				name:     "nil input",
				input:    nil,
				fn:       func(i int) int { return i * 2 },
				expected: nil,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := s.Map(tt.input, tt.fn)
				assert.Equal(t, tt.expected, result)
			})
		}
	}
	{
		tests := []struct {
			name     string
			input    []string
			fn       func(string) string
			expected []string
		}{
			{
				name:     "square each element",
				input:    []string{"a", "b", "c"},
				fn:       func(s string) string { return s + "!" },
				expected: []string{"a!", "b!", "c!"},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := s.Map(tt.input, tt.fn)
				assert.Equal(t, tt.expected, result)
			})
		}
	}
}

func TestMapInPlace(t *testing.T) {
	{
		tests := []struct {
			name     string
			input    []int
			fn       func(int) int
			expected []int
		}{
			{
				name:     "double each element in place",
				input:    []int{1, 2, 3},
				fn:       func(i int) int { return i * 2 },
				expected: []int{2, 4, 6},
			},
			{
				name:     "empty slice in place",
				input:    []int{},
				fn:       func(i int) int { return i * 2 },
				expected: []int{},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				s.MapInPlace(tt.input, tt.fn)
				assert.Equal(t, tt.expected, tt.input)
			})
		}
	}
	{
		tests := []struct {
			name     string
			input    []string
			fn       func(string) string
			expected []string
		}{
			{
				name:     "append '!' to each element in place",
				input:    []string{"a", "b", "c"},
				fn:       func(s string) string { return s + "!" },
				expected: []string{"a!", "b!", "c!"},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				s.MapInPlace(tt.input, tt.fn)
				assert.Equal(t, tt.expected, tt.input)
			})
		}
	}
}
