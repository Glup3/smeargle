package pokemon

import (
	"testing"
)

func TestGetSlugsNoGenerationsCount(t *testing.T) {
	const expectedCount = 905

	config, err := NewPokemonConfig()
	if err != nil {
		t.Errorf("%+v", err)
	}

	slugs := config.GetSlugs()
	if len(slugs) != expectedCount {
		t.Errorf("got %d, want %d", len(slugs), expectedCount)
	}
}
