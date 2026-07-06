package example_country

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestCountry(t *testing.T) {
	assert.NotNil(t, New())
}
