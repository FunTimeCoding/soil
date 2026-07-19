package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/client"
	"testing"
)

func TestClient(t *testing.T) {
	assert.NotNil(t, client.New("", nil))
}
