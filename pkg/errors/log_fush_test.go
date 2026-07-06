package errors

import (
	"github.com/funtimecoding/soil/pkg/errors/mock_flusher"
	"testing"
)

func TestLogFlush(t *testing.T) {
	LogFlush(mock_flusher.New())
}
