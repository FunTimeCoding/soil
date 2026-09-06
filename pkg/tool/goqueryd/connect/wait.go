package connect

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goquery/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"time"
)

func Wait(l *logger.Logger) *client.Client {
	base := locator.Environment(
		constant.HostEnvironment,
		constant.PortEnvironment,
		constant.InsecureEnvironment,
	).String()
	c, e := client.NewClient(
		base,
		client.WithRequestEditorFn(
			web.BearerEditor(environment.Required(constant.TokenEnvironment)),
		),
	)
	errors.PanicOnError(e)
	deadline := time.Now().Add(time.Minute)

	for time.Now().Before(deadline) {
		_, f := c.GetStatus(context.Background())

		if f == nil {
			l.Structured("goqueryd_connected", "address", base)

			return c
		}

		time.Sleep(time.Second)
	}

	panic(fmt.Sprintf("goqueryd not reachable at %s after 1 minute", base))
}
