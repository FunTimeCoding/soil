package runner_option

import (
	"context"
	library "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/provision/face"
	"github.com/funtimecoding/soil/pkg/provision/types/change"
)

type Option struct {
	Repository    string
	ClonePath     string
	ToolPath      string
	ApplyFunction func(
		parameters map[string]any,
		triggerSource string,
	) any
	InitFunction    func()
	SetupFunction   func(context.Context) bool
	CleanupFunction func()
	ChangeFunction  func(value any) *change.Change
	Downstream      []face.Downstream
	Registry        library.ProcessRegistry
}
