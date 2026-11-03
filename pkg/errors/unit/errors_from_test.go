package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	soilErrors "github.com/funtimecoding/soil/pkg/errors"
	"testing"
)

func TestFrom(t *testing.T) {
	assert.Any(t, nil, soilErrors.From(nil))
	assert.Any(t, nil, soilErrors.From(""))
	assert.Any(t, errors.New("a"), soilErrors.From(errors.New("a")))
	assert.Any(t, errors.New("non-error: a"), soilErrors.From("a"))
}
