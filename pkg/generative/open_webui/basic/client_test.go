package basic

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestClient(t *testing.T) {
	assert.NotNil(t, New("", ""))
}
