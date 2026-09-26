package service

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"path/filepath"
	"strings"
)

func (c *Client) launchServices() []*Service {
	r := run.New()
	r.Panic = false
	r.Start(constant.Launchctl, constant.LaunchctlList)
	loaded := ParseLaunchctl(r.OutputString)
	var result []*Service

	for _, directory := range []string{
		constant.LaunchDaemonDirectory,
		constant.LaunchAgentDirectory,
		filepath.Join(c.home, constant.UserLaunchAgents),
	} {
		if !system.DirectoryExists(directory) {
			continue
		}

		for _, name := range system.FilesByExtension(
			directory,
			constant.PropertyListExtension,
		) {
			label := strings.TrimSuffix(
				filepath.Base(name),
				constant.PropertyListExtension,
			)
			current, okay := loaded[label]

			if !okay {
				current = constant.ServiceStopped
			}

			result = append(
				result,
				NewService(
					label,
					current,
					constant.ServiceOriginLocal,
					directory,
					"",
					true,
				),
			)
		}
	}

	return result
}
