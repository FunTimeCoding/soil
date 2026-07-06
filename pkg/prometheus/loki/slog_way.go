package loki

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/grafana/loki-client-go/loki"
	"github.com/samber/slog-loki/v3"
	"log/slog"
)

func slogWay(host string) {
	// Another way to log, not to read
	configuration, e := loki.NewDefaultConfig(
		locator.New(host).Base(constant.Base).Path(constant.Push).String(),
	)
	errors.PanicOnError(e)
	configuration.TenantID = "exampleTenant"
	c, f := loki.New(configuration)
	errors.PanicOnError(f)
	l := slog.New(
		slogloki.Option{Level: slog.LevelDebug, Client: c}.NewLokiHandler(),
	)
	l = l.With("environment", "dev").With("release", "v1.0.0")
	l.Error("example error")
	l.Info("example info")
	c.Stop()
}
