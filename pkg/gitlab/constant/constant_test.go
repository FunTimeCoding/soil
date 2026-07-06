package constant

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "X-Gitlab-Token", TokenHeader)
	assert.Count(t, 3, RequestStates)
}
