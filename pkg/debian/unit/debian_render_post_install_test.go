package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/debian"
	debianConstant "github.com/funtimecoding/soil/pkg/debian/constant"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestRenderPostInstallRestart(t *testing.T) {
	assert.String(
		t,
		`#!/bin/sh
set -e

if [ "$1" != "configure" ]; then
    exit 0
fi

if [ ! -d /run/systemd/system ]; then
    exit 0
fi

systemctl daemon-reload
systemctl restart Alfa.service
`,
		debian.RenderPostInstall(
			constant.UpperAlfa,
			debianConstant.UpgradeRestart,
		),
	)
}

func TestRenderPostInstallKeep(t *testing.T) {
	assert.String(
		t,
		`#!/bin/sh
set -e

if [ "$1" != "configure" ]; then
    exit 0
fi

if [ ! -d /run/systemd/system ]; then
    exit 0
fi

systemctl daemon-reload
`,
		debian.RenderPostInstall(
			constant.UpperAlfa,
			debianConstant.UpgradeKeep,
		),
	)
}
