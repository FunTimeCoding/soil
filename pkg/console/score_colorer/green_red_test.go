package score_colorer

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/console/score_colorer/score_fixture"
	"testing"
)

func TestGreenRed(t *testing.T) {
	assert.NotNil(t, GreenRed(score_fixture.New(1)))
}
