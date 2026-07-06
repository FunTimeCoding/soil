package constant

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "KUBERNETES_CONTEXT", ContextEnvironment)
	assert.String(t, "type=Normal", TypeNormal)
	assert.String(t, "pods", PodsResource)
}
