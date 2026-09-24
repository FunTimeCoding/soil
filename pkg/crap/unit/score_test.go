package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/score"
	"testing"
)

func TestScoreUncovered(t *testing.T) {
	assert.Float(t, 30, score.Score(5, 0))
	assert.Float(t, 72, score.Score(8, 0))
}

func TestScoreCovered(t *testing.T) {
	assert.Float(t, 5, score.Score(5, 100))
}

func TestScorePartial(t *testing.T) {
	assert.Float(t, 5.390625, score.Score(5, 75))
}

func TestParsePolicy(t *testing.T) {
	assert.True(t, score.ParsePolicy("") == constant.Pessimistic)
	assert.True(t, score.ParsePolicy("Optimistic") == constant.Optimistic)
	assert.True(t, score.ParsePolicy("SKIP") == constant.Skip)
}

func TestParsePolicyUnknownPanics(t *testing.T) {
	defer func() { assert.NotNil(t, recover()) }()
	score.ParsePolicy("hopeful")
}
