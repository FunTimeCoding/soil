package service

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/funtimecoding/soil/pkg/system/types/service"
	"path/filepath"
	"strings"
)

func (c *Client) launchServices() []*service.Service {
	r := run.New()
	r.Panic = false
	r.Start(constant.Launchctl, constant.LaunchctlList)
	loaded := ParseLaunchctl(r.OutputString)
	var result []*service.Service

	for _, directory := range []string{
		constant.LaunchDaemonDirectory,
		constant.LaunchAgentDirectory,
		filepath.Join(system.Home(), constant.UserLaunchAgents),
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
				service.New(
					label,
					current,
					constant.ServiceOriginLocal,
					directory,
					"",
					"",
					true,
				),
			)
		}
	}

	return result
}
