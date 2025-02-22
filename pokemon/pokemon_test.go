package pokemon

import (
	"fmt"
	"testing"
)

func TestGetSlugsCount(t *testing.T) {
	tests := []struct {
		gens  []int
		count int
	}{
		{[]int{}, 905},
		{[]int{1}, 151},
		{[]int{2}, 100},
		{[]int{3}, 135},
		{[]int{4}, 107},
		{[]int{5}, 156},
		{[]int{6}, 72},
		{[]int{7}, 88},
		{[]int{8}, 96},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 905},
		{[]int{1, 3}, 286},
		{[]int{2, 4, 5}, 363},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%+v", tt.gens)
		t.Run(testname, func(t *testing.T) {
			config, err := NewPokemonConfig()
			if err != nil {
				t.Errorf("%+v", err)
			}

			slugs, err := config.GetSlugs(tt.gens)
			if len(slugs) != tt.count {
				t.Errorf("got %d, want %d", len(slugs), tt.count)
			}
		})
	}
}
