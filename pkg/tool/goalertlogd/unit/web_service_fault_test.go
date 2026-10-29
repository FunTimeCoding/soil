package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/web_service_tester"
	"net/http"
	"testing"
)

func TestAStoreFaultCarriesItsOwnText(t *testing.T) {
	o := web_service_tester.New(t)
	o.Service().Close()
	r, e := o.Client.GetAlertsWithResponse(
		context.Background(),
		&client.GetAlertsParams{Name: "HighMemory"},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusInternalServerError, r.StatusCode())
	assert.String(t, "sql: database is closed", r.JSON500.Error)
}
