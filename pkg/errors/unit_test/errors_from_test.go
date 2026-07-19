package unit_test

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	library "github.com/funtimecoding/soil/pkg/errors"
	"testing"
)

func TestFrom(t *testing.T) {
	assert.Any(t, nil, library.From(nil))
	assert.Any(t, nil, library.From(""))
	assert.Any(t, errors.New("a"), library.From(errors.New("a")))
	assert.Any(t, errors.New("non-error: a"), library.From("a"))
}
