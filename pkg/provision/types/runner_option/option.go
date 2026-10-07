package runner_option

import "github.com/funtimecoding/soil/pkg/face"

type Option struct {
	Repository    string
	ClonePath     string
	ToolPath      string
	ApplyFunction func(
		parameters map[string]any,
		triggerSource string,
	) any
	InitFunction    func()
	SetupFunction   func() bool
	CleanupFunction func()
	ChangeFunction  func(value any) []string
	Downstream      []face.Downstream
	Registry        face.ProcessRegistry
}
