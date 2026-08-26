package debian

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/debian/constant"
)

func RenderPreRemove(executableName string) string {
	return fmt.Sprintf(
		`#!/bin/sh
set -e

if [ "$1" != "remove" ]; then
    exit 0
fi

if [ ! -d /run/systemd/system ]; then
    exit 0
fi

systemctl stop %s.%s
systemctl disable %s.%s
`,
		executableName,
		constant.ServiceExtension,
		executableName,
		constant.ServiceExtension,
	)
}
