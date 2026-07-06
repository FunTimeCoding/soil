package age_colorer

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/console/age_colorer/age_fixture"
	"testing"
	"time"
)

func TestGreenRed(t *testing.T) {
	assert.NotNil(t, GreenRed(age_fixture.New(0*time.Hour)))
}
