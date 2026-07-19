package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/console/age_colorer"
	"github.com/funtimecoding/soil/pkg/console/age_colorer/age_fixture"
	"testing"
	"time"
)

func TestAgeColorerGreenRed(t *testing.T) {
	assert.NotNil(t, age_colorer.GreenRed(age_fixture.New(0*time.Hour)))
}
