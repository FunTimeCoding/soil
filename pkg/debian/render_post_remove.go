package debian

func RenderPostRemove() string {
	return `#!/bin/sh
set -e

if [ "$1" != "remove" ] && [ "$1" != "purge" ]; then
    exit 0
fi

if [ ! -d /run/systemd/system ]; then
    exit 0
fi

systemctl daemon-reload
`
}
