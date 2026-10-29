package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/client"
)

type Client interface {
	face.Reachable
	Address() string
	Processes(x context.Context) ([]client.Process, error)
	Start(
		x context.Context,
		name string,
	) error
	Stop(
		x context.Context,
		name string,
	) error
	Restart(
		x context.Context,
		name string,
	) error
	RestartAll(x context.Context) (int, error)
	ReloadProcfile(x context.Context) error
	ReloadEnvironment(x context.Context) error
	Log(
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
