package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"testing"
)

func TestMarkResponse(t *testing.T) {
	_, e := response.Fail("")
	assert.Nil(t, e)
	_, f := response.Success("")
	assert.Nil(t, f)
	_, g := response.FailAny("")
	assert.Nil(t, g)
	_, h := response.SuccessAny("")
	assert.Nil(t, h)
}
