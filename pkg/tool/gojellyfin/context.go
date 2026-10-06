package gojellyfin

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
)

type Context struct {
	Client    *client.ClientWithResponses
	Telemetry face.Recorder
	Terminal  *terminal.Terminal
}
