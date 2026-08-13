package example

import (
	"github.com/akkuman/logrus-loki-hook"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/sirupsen/logrus"
)

func Write() {
	host := environment.Required(constant.LokiHostEnvironment)
	user := environment.Required(constant.LokiUserEnvironment)
	password := environment.Required(constant.LokiPasswordEnvironment)
	l := logrus.New()
	h, e := hook.NewHook(
		&hook.Config{
			URL: locator.New(host).UserPassword(user, password).Path(
				"/api/prom/push",
			).String(),
			LevelName: "severity",
			Labels:    map[string]string{"application": "test"},
		},
	)

	if e != nil {
		l.Error(e)
	} else {
		l.AddHook(h)
	}

	l.Info("test message")
}
