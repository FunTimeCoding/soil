package response

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestResponse(t *testing.T) {
	_, e := Fail("")
	assert.Nil(t, e)
	_, f := Success("")
	assert.Nil(t, f)
	_, g := FailAny("")
	assert.Nil(t, g)
	_, h := SuccessAny("")
	assert.Nil(t, h)
}
