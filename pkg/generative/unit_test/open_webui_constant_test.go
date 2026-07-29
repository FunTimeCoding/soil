package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"testing"
)

func TestOpenWebuiConstant(t *testing.T) {
	assert.String(t, "OPEN_WEBUI_HOST", constant.OpenWebInterfaceHostEnvironment)
	assert.String(
		t,
		"OPEN_WEBUI_TOKEN",
		constant.OpenWebInterfaceTokenEnvironment,
	)
}
