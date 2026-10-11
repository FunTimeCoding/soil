package debian

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/constant"
)

func RenderPostInstall(
	executableName string,
	upgradeMode string,
) string {
	restart := ""

	if upgradeMode == constant.DebianUpgradeRestart {
		restart = fmt.Sprintf(
			"systemctl restart %s.%s\n",
			executableName,
			constant.DebianServiceExtension,
		)
	}

	return fmt.Sprintf(
		`#!/bin/sh
set -e

if [ "$1" != "configure" ]; then
    exit 0
fi

if [ ! -d /run/systemd/system ]; then
    exit 0
fi

systemctl daemon-reload
%s`,
		restart,
	)
}
