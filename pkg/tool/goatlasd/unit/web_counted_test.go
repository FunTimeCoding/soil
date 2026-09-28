package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/web"
	"testing"
)

func TestCountedSpellsOneSingular(t *testing.T) {
	assert.String(
		t,
		"1 placement",
		web.Counted(1, constant.PlacementWord, constant.PlacementsWord),
	)
}

func TestCountedSpellsEverythingElsePlural(t *testing.T) {
	assert.String(
		t,
		"0 sightings",
		web.Counted(0, constant.SightingWord, constant.SightingsWord),
	)
	assert.String(
		t,
		"52 sightings",
		web.Counted(52, constant.SightingWord, constant.SightingsWord),
	)
}

func TestWordCarriesNoCount(t *testing.T) {
	assert.String(
		t,
		"place",
		web.Word(1, constant.PlaceWord, constant.PlacesWord),
	)
	assert.String(
		t,
		"places",
		web.Word(4, constant.PlaceWord, constant.PlacesWord),
	)
}
