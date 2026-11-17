package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/base"
	"net/http"
	"testing"
)

func TestStoreFaultCarriesItsOwnText(t *testing.T) {
	s := base.New(t)
	c := s.GeneratedClient()
	s.Daemon.Close()
	r, e := c.GetAlertsWithResponse(
		context.Background(),
		&client.GetAlertsParams{Name: "HighMemory"},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusInternalServerError, r.StatusCode())
	assert.String(t, "sql: database is closed", r.JSON500.Error)
}
