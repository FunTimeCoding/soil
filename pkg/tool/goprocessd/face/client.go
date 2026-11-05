package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/client"
)

type Client interface {
	face.Reachable
	Address() string
	ListProcesses(x context.Context) ([]client.Process, error)
	StartProcess(
		x context.Context,
		name string,
	) error
	StopProcess(
		x context.Context,
		name string,
	) error
	RestartProcess(
		x context.Context,
		name string,
	) error
	RestartAll(x context.Context) (int, error)
	ReloadProcfile(x context.Context) error
	ReloadEnvironment(x context.Context) error
	GetLog(
		x context.Context,
		name string,
		all bool,
	) (*client.Log, error)
	ClearLog(
		x context.Context,
		name string,
	) error
	Close()
}
