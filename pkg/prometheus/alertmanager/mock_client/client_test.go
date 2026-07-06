package mock_client

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestClient(t *testing.T) {
	assert.NotNil(t, New())
}
