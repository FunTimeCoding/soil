package monitor

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/claim"
	"github.com/funtimecoding/soil/pkg/bubbletea/table/item"
	"github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/client"
	monitorConstant "github.com/funtimecoding/soil/pkg/tool/gomonitord/constant"
)

func New(connect bool) *Model {
	result := &Model{
		table:    item.New(connect),
		connect:  connect,
		claims:   map[string][]string{},
		user:     system.User().Username,
		hostname: system.Hostname(),
		auto:     !environment.Exists(constant.ManualEnvironment),
	}
	result.owner = fmt.Sprintf(
		constant.OwnerFormat,
		result.user,
		result.hostname,
	)

	if !connect {
		return result
	}

	if !environment.Exists(monitorConstant.TokenEnvironment) {
		result.notice = fmt.Sprintf(
			"%s is not set - claims are off",
			monitorConstant.TokenEnvironment,
		)

		return result
	}

	result.monitor = client.NewEnvironment()
	result.updates = make(chan claim.Message, 1)

	return result
}
