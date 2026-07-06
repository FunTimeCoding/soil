package chat_request

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestRequest(t *testing.T) {
	assert.NotNil(t, New())
}
