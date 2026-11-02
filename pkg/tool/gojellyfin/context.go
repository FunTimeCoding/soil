package gojellyfin

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/terminal"
	jellyfin "github.com/funtimecoding/soil/pkg/tool/gojellyfind/face"
)

type Context struct {
	Client    jellyfin.Client
	Telemetry face.Recorder
	Terminal  *terminal.Terminal
}
