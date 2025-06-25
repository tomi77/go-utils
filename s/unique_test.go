package s_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomi77/go-utils/s"
)

func TestUnique(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "unique elements",
			input:    []int{1, 2, 3, 1, 2},
			expected: []int{1, 2, 3},
		},
		{
			name:     "all unique",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "nil slice",
			input:    nil,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := s.Unique(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestUniqueFunc(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		fn       func(string) string
		expected []string
	}{
		{
			name:     "unique transformed elements",
			input:    []string{"a", "b", "c", "a"},
			fn:       func(s string) string { return s },
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "transform to unique elements",
			input:    []string{"apple", "banana", "applE", "orange"},
			fn:       func(s string) string { return s[:3] }, // transform to first 3 characters
			expected: []string{"apple", "banana", "orange"},
		},
		{
			name:     "empty slice",
			input:    []string{},
			fn:       func(s string) string { return s },
			expected: []string{},
		},
		{
			name:     "nil slice",
			input:    nil,
			fn:       func(s string) string { return s },
			expected: nil,
		},
		{
			name:     "nil function",
			input:    []string{"a", "b", "c"},
			fn:       nil,
			expected: nil,
		},
		{
			name:     "case insensitive uniqueness",
			input:    []string{"apple", "Apple", "banana"},
			fn:       strings.ToLower,
			expected: []string{"apple", "banana"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := s.UniqueFunc(tt.input, tt.fn)
			assert.Equal(t, tt.expected, result)
		})
	}
}
