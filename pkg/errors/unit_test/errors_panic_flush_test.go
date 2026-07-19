package unit_test

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/mock_flusher"
	"testing"
)

func TestPanicFlush(t *testing.T) {
	errors.PanicFlush(mock_flusher.New())
}
