package goatlas

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/generated/client"
)

type Context struct {
	Client    *client.ClientWithResponses
	Telemetry face.Recorder
}
