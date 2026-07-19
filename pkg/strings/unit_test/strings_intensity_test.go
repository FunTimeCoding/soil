package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestIntensity(t *testing.T) {
	v := []string{upper.Alfa, upper.Bravo, upper.Charlie}
	assert.String(t, "Alfa", strings.Intensity(v, 0))
	assert.String(t, "Alfa", strings.Intensity(v, 0.33))
	assert.String(t, "Bravo", strings.Intensity(v, 0.34))
	assert.String(t, "Bravo", strings.Intensity(v, 0.66))
	assert.String(t, "Charlie", strings.Intensity(v, 0.67))
	assert.String(t, "Charlie", strings.Intensity(v, 1))
}
