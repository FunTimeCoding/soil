package connect

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goquery/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"time"
)

func Wait(l *logger.Logger) *client.Client {
	host := environment.Required(constant.HostEnvironment)
	port := environment.RequiredInteger(constant.PortEnvironment)
	base := locator.New(host).Port(port).Insecure().String()
	c, e := client.NewClient(base)
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
