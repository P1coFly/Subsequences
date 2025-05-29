package main

import "testing"

func TestMinAlphabetWindowLength(t *testing.T) {
	tests := []struct {
		name     string
		seq      []int
		expected int
	}{
		{
			name:     "Too short, can't contain all 26 letters",
			seq:      []int{1, 10, 6, 7, 18},
			expected: 0,
		},
		{
			name: "Contains all 26 letters",
			seq: []int{
				1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
				11, 12, 13, 14, 15, 16, 17, 18,
				19, 20, 22, 21, 23, 25, 26, 1, 24, 1,
			},
			expected: 26,
		},
		{
			name:     "Exactly 26 in order",
			seq:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26},
			expected: 26,
		},
		{
			name:     "All elements are the same",
			seq:      []int{1, 1, 1, 1, 1, 1},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minAlphabetWindowLength(tt.seq)
			if result != tt.expected {
				t.Errorf("got %d, expected %d", result, tt.expected)
			}
		})
	}
}
