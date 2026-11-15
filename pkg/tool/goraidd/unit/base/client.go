package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/client"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/option"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func Client(
	t *testing.T,
	o *option.Raid,
	token string,
) *client.Client {
	t.Helper()
	result := client.New(
		locator.New("localhost").Port(
			Mount(t, o, memory.New()).Port,
		).Insecure(),
		token,
	)
	t.Cleanup(result.Close)

	return result
}
