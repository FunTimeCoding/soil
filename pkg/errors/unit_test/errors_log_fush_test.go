package unit_test

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/mock_flusher"
	"testing"
)

func TestLogFlush(t *testing.T) {
	errors.LogFlush(mock_flusher.New())
}
