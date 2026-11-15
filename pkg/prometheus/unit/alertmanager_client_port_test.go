package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/unit/alertmanager_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"testing"
)

func TestAlertmanagerClientDialsConfiguredPort(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(`{}`),
		upstream_server.WithObject(),
	)
	_, e := alertmanager_tester.New(t, s.URL).Status()
	assert.FatalOnError(t, e)
	assert.String(t, "/api/v2/status", s.Last().URL.Path)
}
