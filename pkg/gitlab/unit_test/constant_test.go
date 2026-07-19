package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "X-Gitlab-Token", constant.TokenHeader)
	assert.Count(t, 3, constant.RequestStates)
}
