package store_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/generative/ollama/reachable_skip"
)

func (o *Tester) client() *ollama.Client {
	o.t.Helper()
	result := ollama.NewEnvironment()
	o.t.Cleanup(result.Close)
	reachable_skip.Skip(o.t)

	return result
}
