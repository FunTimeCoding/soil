package connect

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/system/environment"
	memoryClient "github.com/funtimecoding/soil/pkg/tool/gomemoryd/client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	generated "github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"time"
)

func Wait(l *logger.Logger) memoryClient.Client {
	host := environment.Required(constant.HostEnvironment)
	port := environment.RequiredInteger(constant.PortEnvironment)
	base := locator.New(host).Port(port).Insecure().String()
	c, e := generated.NewClient(base)
	errors.PanicOnError(e)
	deadline := time.Now().Add(time.Minute)

	for time.Now().Before(deadline) {
		r, f := c.GetVersions(
			context.Background(),
			&generated.GetVersionsParams{Since: "2099-01-01T00:00:00Z"},
		)

		if f == nil {
			errors.PanicClose(r.Body)
			l.Structured("gomemoryd_connected", "address", base)

			return memoryClient.NewRestClient(c)
		}

		time.Sleep(time.Second)
	}

	panic(fmt.Sprintf("gomemoryd not reachable at %s after 1 minute", base))
}
