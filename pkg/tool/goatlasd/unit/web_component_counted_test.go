package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/web/component"
	"testing"
)

func TestCountedSpellsOneSingular(t *testing.T) {
	assert.String(
		t,
		"1 placement",
		component.Counted(1, constant.PlacementWord, constant.PlacementsWord),
	)
}

func TestCountedSpellsEverythingElsePlural(t *testing.T) {
	assert.String(
		t,
		"0 sightings",
		component.Counted(0, constant.SightingWord, constant.SightingsWord),
	)
	assert.String(
		t,
		"52 sightings",
		component.Counted(52, constant.SightingWord, constant.SightingsWord),
	)
}

func TestWordCarriesNoCount(t *testing.T) {
	assert.String(
		t,
		"place",
		component.Word(1, constant.PlaceWord, constant.PlacesWord),
	)
	assert.String(
		t,
		"places",
		component.Word(4, constant.PlaceWord, constant.PlacesWord),
	)
}
