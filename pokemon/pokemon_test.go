package pokemon

import (
	"fmt"
	"testing"
)

func TestGetSlugsCount(t *testing.T) {
	tests := []struct {
		gens          []int
		expectedCount int
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
		{[]int{9}, 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("gens %+v", tt.gens)
		t.Run(testname, func(t *testing.T) {
			config, err := NewPokemonConfig()
			if err != nil {
				t.Errorf("%+v", err)
			}

			slugs, err := config.GetSlugs(tt.gens, Idx, Asc)
			if len(slugs) != tt.expectedCount {
				t.Errorf("got %d, expected %d", len(slugs), tt.expectedCount)
			}
		})
	}
}

func TestGetSlugsSorting(t *testing.T) {
	tests := []struct {
		gens          []int
		orderBy       OrderBy
		sortDir       SortDirection
		expectedFirst string
		expectedLast  string
	}{
		{[]int{}, Alphabet, Asc, "abomasnow", "zygarde"},
		{[]int{}, Alphabet, Desc, "zygarde", "abomasnow"},
		{[]int{}, Idx, Asc, "bulbasaur", "enamorus"},
		{[]int{}, Idx, Desc, "enamorus", "bulbasaur"},
		{[]int{2}, Alphabet, Asc, "aipom", "yanma"},
		{[]int{2}, Alphabet, Desc, "yanma", "aipom"},
		{[]int{2}, Idx, Asc, "chikorita", "celebi"},
		{[]int{2}, Idx, Desc, "celebi", "chikorita"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("gens %+v by %v %v", tt.gens, tt.orderBy, tt.sortDir)
		t.Run(testname, func(t *testing.T) {
			config, err := NewPokemonConfig()
			if err != nil {
				t.Errorf("%+v", err)
			}

			slugs, err := config.GetSlugs(tt.gens, tt.orderBy, tt.sortDir)
			if slugs[0] != tt.expectedFirst {
				t.Errorf("got %s, expected %s", slugs[0], tt.expectedFirst)
			}
			if slugs[len(slugs)-1] != tt.expectedLast {
				t.Errorf("got %s, expected %s", slugs[len(slugs)-1], tt.expectedLast)
			}
		})
	}
}
