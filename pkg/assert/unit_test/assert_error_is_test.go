package unit_test

import (
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestErrorIs(t *testing.T) {
	sentinel := errors.New("not found")
	assert.ErrorIs(t, fmt.Errorf("thing not found: %w", sentinel), sentinel)
}
