package gohabitica

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/client"
)

type Context struct {
	Client    *client.Client
	Telemetry face.Recorder
}
