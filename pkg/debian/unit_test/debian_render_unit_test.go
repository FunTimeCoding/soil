package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/debian"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestRenderUnit(t *testing.T) {
	assert.String(
		t,
		`[Unit]
Description=Alfa stub description
After=network.target

[Service]
Type=simple
ExecStart=/usr/local/bin/Alfa

[Install]
WantedBy=multi-user.target
`,
		debian.RenderUnit(constant.UpperAlfa),
	)
}
