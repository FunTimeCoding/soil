package unit

import (
	"errors"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/assert"
	soilErrors "github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"testing"
)

func TestDeleteWhoseRefreshFailsStillAnswersAndReports(t *testing.T) {
	r := memory.New()
	d, clients := daemon_tester.New(t, r, constant.FixtureInstance)
	clients[constant.FixtureInstance].FailProjects(errors.New("unreachable"))
	v := startServer(t, d, r)
	defer v.Close()
	q := web.NewPost(
		locator.New(webConstant.Localhost).Insecure().Port(v.Port).Path(
			constant.DeletePath,
		).Set(
			inventoryConstant.InstanceParameter,
			constant.FixtureInstance,
		).SetInteger(argumentConstant.Project, 1).SetInteger(
			argumentConstant.Pipeline,
			10,
		).String(),
		"",
	)
	o, e := v.Client().Do(q)
	assert.FatalOnError(t, e)
	soilErrors.PanicClose(o.Body)
	assert.Integer(t, http.StatusOK, o.StatusCode)
	assert.String(t, "/", o.Header.Get(webConstant.ExtendedRedirect))
	assert.Count(t, 1, r.Events())
}
