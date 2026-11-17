package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/console/score_scale"
	"github.com/funtimecoding/soil/pkg/console/score_scale/score_fixture"
	"testing"
)

func TestScoreScaleGreenRed(t *testing.T) {
	assert.NotNil(t, score_scale.GreenRed(score_fixture.New(1)))
}
