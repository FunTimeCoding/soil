package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "Accept-Language", constant.AcceptLanguage)
	assert.String(t, "User-Agent", constant.UserAgent)
	assert.String(t, "image/x-icon", constant.Icon)
	assert.String(t, "method", constant.FormMethod)
}
