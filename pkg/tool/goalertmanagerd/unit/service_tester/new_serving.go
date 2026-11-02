package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/face"
	prometheusFace "github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/option"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/service"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewServing(
	t *testing.T,
	payload string,
) *service.Service {
	t.Helper()
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				_ *http.Request,
			) {
				w.Header().Set(webConstant.ContentType, webConstant.Object)
				_, e := w.Write([]byte(payload))
				errors.PanicOnError(e)
			},
		),
	)
	t.Cleanup(s.Close)
	i := instance.Instance{Name: "alfa"}
	i.Alertmanager.Host = "127.0.0.1"
	i.Alertmanager.Port = s.Listener.Addr().(*net.TCPAddr).Port
	o := option.New()
	o.Inventory = inventory.New[instance.Instance](i)
	o.Alertmanager = func(i *instance.Instance) (face.Client, error) {
		return alertmanager.New(
			i.Alertmanager.Host,
			i.Alertmanager.Port,
			false,
			"",
			"",
		), nil
	}
	o.Prometheus = func(*instance.Instance) (prometheusFace.Client, error) {
		return mock_client.New(), nil
	}
	result := service.New(o, logger.New(t.Context()), memory.New())
	t.Cleanup(result.Close)

	return result
}
