package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/console/age_scale"
	"github.com/funtimecoding/soil/pkg/console/age_scale/age_fixture"
	"testing"
	"time"
)

func TestAgeScaleGreenRed(t *testing.T) {
	assert.NotNil(t, age_scale.GreenRed(age_fixture.New(0*time.Hour)))
}
