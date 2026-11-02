package unit

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/daemon_tester"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"testing"
)

func TestWebRefusesAnUnresolvableInstanceAsBadRequest(t *testing.T) {
	r := memory.New()
	d, _ := daemon_tester.New(
		t,
		r,
		constant.FixtureInstance,
		constant.FixtureOtherInstance,
	)
	v := startServer(t, d, r)
	defer v.Close()
	unknown := locator.New(webConstant.Localhost).Insecure().Port(v.Port).Path(
		constant.PipelinePath,
	).Set(inventoryConstant.InstanceParameter, "charlie").SetInteger(
		argumentConstant.Project,
		1,
	).SetInteger(argumentConstant.Pipeline, 1)
	assert.HTTPStatus(t, v.Client(), unknown.String(), http.StatusBadRequest)
	missing := locator.New(webConstant.Localhost).Insecure().Port(v.Port).Path(
		constant.PipelinePath,
	).SetInteger(argumentConstant.Project, 1).SetInteger(
		argumentConstant.Pipeline,
		1,
	)
	assert.HTTPStatus(t, v.Client(), missing.String(), http.StatusBadRequest)
}
