package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/debian"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestRenderPreRemove(t *testing.T) {
	assert.String(
		t,
		`#!/bin/sh
set -e

if [ "$1" != "remove" ]; then
    exit 0
fi

if [ ! -d /run/systemd/system ]; then
    exit 0
fi

systemctl stop Alfa.service
systemctl disable Alfa.service
`,
		debian.RenderPreRemove(constant.UpperAlfa),
	)
}

func TestRenderPostRemove(t *testing.T) {
	assert.String(
		t,
		`#!/bin/sh
set -e

if [ "$1" != "remove" ] && [ "$1" != "purge" ]; then
    exit 0
fi

if [ ! -d /run/systemd/system ]; then
    exit 0
fi

systemctl daemon-reload
`,
		debian.RenderPostRemove(),
	)
}
