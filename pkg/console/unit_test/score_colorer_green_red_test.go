package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/console/score_colorer"
	"github.com/funtimecoding/soil/pkg/console/score_colorer/score_fixture"
	"testing"
)

func TestScoreColorerGreenRed(t *testing.T) {
	assert.NotNil(t, score_colorer.GreenRed(score_fixture.New(1)))
}
