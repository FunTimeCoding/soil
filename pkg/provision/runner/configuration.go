package runner

import "github.com/funtimecoding/soil/pkg/face"

type Configuration struct {
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
	Registry        face.ProcessRegistry
}
