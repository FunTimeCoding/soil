package runner

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/types/host_command"
	"strings"
)

func (r *Runner) hostPattern(
	d *host_command.HostCommand,
	host string,
	directory string,
	inventoryDirectory string,
) (string, error) {
	if host == "" {
		return d.Group, nil
	}

	pattern := fmt.Sprintf(constant.HostCommandPattern, d.Group, host)
	list := r.newRun().NoPanic()
	list.Directory = directory
	list.Start(adHoc(inventoryDirectory, pattern, constant.ListHostsFlag)...)

	if list.Error != nil {
		return "", fmt.Errorf(
			constant.ListHostsFailed,
			list.Error,
			strings.TrimSpace(list.ErrorString),
		)
	}

	if !listed(list.OutputString, host) {
		return "", validation.New(constant.HostCommandOutside, host, d.Group)
	}

	return pattern, nil
}
