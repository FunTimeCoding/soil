package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/open_webui/basic"
	"testing"
)

func TestClient(t *testing.T) {
	assert.NotNil(t, basic.New("", ""))
}
