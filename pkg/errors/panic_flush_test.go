package errors

import (
	"github.com/funtimecoding/soil/pkg/errors/mock_flusher"
	"testing"
)

func TestPanicFlush(t *testing.T) {
	PanicFlush(mock_flusher.New())
}
