package constant

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "Accept-Language", AcceptLanguage)
	assert.String(t, "User-Agent", UserAgent)
	assert.String(t, "image/x-icon", Icon)
	assert.String(t, "method", FormMethod)
}
