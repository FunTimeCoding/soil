package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/prometheus/mock_client"
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

	return pooled(
		t,
		[]instance.Instance{i},
		func(i *instance.Instance) face.AlertmanagerClient {
			return alertmanager.New(
				i.Alertmanager.Host,
				i.Alertmanager.Port,
				false,
				"",
				"",
			)
		},
		func(*instance.Instance) face.PrometheusClient {
			return mock_client.New()
		},
	)
}
