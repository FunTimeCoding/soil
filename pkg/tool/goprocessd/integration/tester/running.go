package tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (t *Tester) Running() string {
	v, e := t.Client.Processes(context.Background())
	errors.PanicOnError(e)
	var result []string

	for _, p := range v {
		if p.Running {
			result = append(result, p.Name)
		}
	}

	return join.Comma(result)
}
