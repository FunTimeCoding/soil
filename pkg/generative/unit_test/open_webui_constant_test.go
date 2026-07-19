package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/open_webui/constant"
	"testing"
)

func TestOpenWebuiConstant(t *testing.T) {
	assert.String(t, "OPEN_WEBUI_HOST", constant.HostEnvironment)
	assert.String(t, "OPEN_WEBUI_TOKEN", constant.TokenEnvironment)
}
