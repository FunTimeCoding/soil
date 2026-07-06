package errors

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestFrom(t *testing.T) {
	assert.Any(t, nil, From(nil))
	assert.Any(t, nil, From(""))
	assert.Any(t, errors.New("a"), From(errors.New("a")))
	assert.Any(t, errors.New("non-error: a"), From("a"))
}
