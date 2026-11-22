package runner_option

import (
	"context"
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/provision/face"
	"github.com/funtimecoding/soil/pkg/provision/types/change"
	"github.com/funtimecoding/soil/pkg/provision/types/command_run"
)

type Option struct {
	Repository     string
	CloneDirectory string
	CopyModes      bool
	ToolPath       string
	ApplyFunction  func(
		parameters map[string]any,
		triggerSource string,
	) any
	InitFunction        func()
	SetupFunction       func(context.Context) bool
	CleanupFunction     func()
	ChangeFunction      func(value any) *change.Change
	HostCommandFunction func(
		command string,
		host string,
		source string,
	) (*command_run.Run, error)
	Downstream []face.Downstream
	Registry   soilFace.ProcessRegistry
}
