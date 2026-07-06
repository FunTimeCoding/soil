package monitor

import (
	"github.com/funtimecoding/soil/pkg/bubbletea/table/item"
	"github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/gorilla/client"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func New(connect bool) *Model {
	return &Model{
		table:    item.New(connect),
		client:   client.New(),
		connect:  connect,
		user:     system.User().Username,
		hostname: system.Hostname(),
		auto:     !environment.Exists(constant.ManualEnvironment),
	}
}
