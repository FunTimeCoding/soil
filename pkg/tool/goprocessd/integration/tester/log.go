package tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/client"
)

func (t *Tester) Log(
	name string,
	all bool,
) *client.Log {
	result, e := t.Client.Log(context.Background(), name, all)
	errors.PanicOnError(e)

	return result
}
