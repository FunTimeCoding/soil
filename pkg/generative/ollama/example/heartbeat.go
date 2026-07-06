package example

import "github.com/funtimecoding/soil/pkg/generative/ollama"

func Heartbeat() {
	o := ollama.NewEnvironment()
	o.MustHeartbeat()
}
