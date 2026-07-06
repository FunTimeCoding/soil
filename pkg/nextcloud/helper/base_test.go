package helper

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestBase(t *testing.T) {
	assert.String(t, "https://example.org", Base(constant.Example))
}
