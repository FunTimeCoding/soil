package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/distance"
	"testing"
)

func TestDistanceLevenshtein(t *testing.T) {
	assert.Integer(t, 3, distance.Levenshtein("kitten", "sitting"))
	assert.Integer(t, 0, distance.Levenshtein("usage", "usage"))
	assert.Integer(t, 5, distance.Levenshtein("", "usage"))
	assert.Integer(t, 5, distance.Levenshtein("usage", ""))
}
