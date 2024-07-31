package main

import (
	"testing"
)

func TestOrder(t *testing.T) {
	tests := []struct {
		name          string
		startingFloor int
		queue         []Person
		expected      []int
	}{
		{
			name:          "Test Case 1",
			startingFloor: 1,
			queue: []Person{
				{From: 3, To: 2}, // Al
				{From: 5, To: 2}, // Betty
				{From: 2, To: 1}, // Charles
				{From: 2, To: 5}, // Dan
				{From: 4, To: 3}, // Ed
			},
			expected: []int{2, 5, 4, 3, 2, 1},
		},
		{
			name:          "Test Case 2",
			startingFloor: 5,
			queue: []Person{
				{From: 5, To: 4},  // 1st passenger
				{From: 5, To: 3},  // 2nd passenger
				{From: 3, To: 4},  // 3rd passenger
				{From: 0, To: 2},  // 5th passenger
				{From: 3, To: -4}, // 4th passenger
				{From: 1, To: 2},
			},
			expected: []int{5, 4, 3, 4, 3, -4, 0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Order(tt.startingFloor, tt.queue)
			if len(result) != len(tt.expected) {
				t.Errorf("Order() = %v, want %v", result, tt.expected)
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Order() = %v, want %v", result, tt.expected)
					return
				}
			}
		})
	}
}
