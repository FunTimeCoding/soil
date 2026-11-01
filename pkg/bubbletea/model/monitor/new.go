package monitor

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/claim"
	"github.com/funtimecoding/soil/pkg/bubbletea/table/item"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/peer/hold"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/client"
	gomonitordConstant "github.com/funtimecoding/soil/pkg/tool/gomonitord/constant"
)

func New(
	connect bool,
	r face.Reporter,
) *Model {
	result := &Model{
		table:    item.New(connect),
		connect:  connect,
		claims:   map[string][]string{},
		user:     system.User().Username,
		hostname: system.Hostname(),
		auto:     !environment.Exists(monitorConstant.ManualEnvironment),
	}
	result.owner = fmt.Sprintf(
		monitorConstant.OwnerFormat,
		result.user,
		result.hostname,
	)

	if !connect {
		return result
	}

	if !environment.Exists(gomonitordConstant.TokenEnvironment) {
		result.notice = fmt.Sprintf(
			"%s is not set - claims are off",
			gomonitordConstant.TokenEnvironment,
		)

		return result
	}

	result.monitor = client.NewEnvironment()
	result.updates = make(chan claim.Message, 1)
	c := claim.New(result.monitor, result.updates)
	result.hold = hold.New(
		gomonitordConstant.Identity.Name(),
		c,
		logger.New(context.Background()),
		r,
	).WithReport(c.Crossed, c.Refused)

	return result
}
