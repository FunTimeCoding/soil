package unit_test

import (
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestError(t *testing.T) {
	assert.Error(t, fmt.Errorf("something went wrong"))
}

func TestErrorIs(t *testing.T) {
	sentinel := errors.New("not found")
	assert.ErrorIs(t, fmt.Errorf("thing not found: %w", sentinel), sentinel)
}

func TestFatalOnError(t *testing.T) {
	assert.FatalOnError(t, nil)
}
