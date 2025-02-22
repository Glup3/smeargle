package pokemon

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseGenerationString(t *testing.T) {
	tests := []struct {
		s            string
		expectedGens []int
	}{
		{"1", []int{1}},
		{"2", []int{2}},
		{"3", []int{3}},
		{"4", []int{4}},
		{"5", []int{5}},
		{"6", []int{6}},
		{"7", []int{7}},
		{"8", []int{8}},
		{"2-5", []int{2, 3, 4, 5}},
		{"1,3", []int{1, 3}},
		{"1,3-5", []int{1, 3, 4, 5}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.s)
		t.Run(testname, func(t *testing.T) {
			gens, err := ParseGenerationString(tt.s)
			if err != nil {
				t.Errorf("%+v", err)
			}

			if !reflect.DeepEqual(gens, tt.expectedGens) {
				t.Errorf("got %+v, expected %+v", gens, tt.expectedGens)
			}
		})
	}
}

func TestParseGenerationStringInvalids(t *testing.T) {
	tests := []struct {
		s string
	}{
		{"-1"},
		{"1-"},
		{"0"},
		{"9"},
		{"0-2"},
		{"-1-3"},
		{"1-3-"},
		{"1-9"},
		{"--"},
		{"ditto"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.s)
		t.Run(testname, func(t *testing.T) {
			_, err := ParseGenerationString(tt.s)
			if err == nil {
				t.Errorf("expected error for: %s", tt.s)
			}
		})
	}
}
