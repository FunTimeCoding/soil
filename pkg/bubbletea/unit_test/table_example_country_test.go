package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/bubbletea/table/example_country"
	"testing"
)

func TestCountry(t *testing.T) {
	assert.NotNil(t, example_country.New())
}
