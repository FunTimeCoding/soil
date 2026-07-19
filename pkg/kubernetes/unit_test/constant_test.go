package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "KUBERNETES_CONTEXT", constant.ContextEnvironment)
	assert.String(t, "type=Normal", constant.TypeNormal)
	assert.String(t, "pods", constant.PodsResource)
}
