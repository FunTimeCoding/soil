package unit

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/assert"
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/service_tester"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"testing"
)

func TestWebRefusesUnresolvableInstanceAsBadRequest(t *testing.T) {
	v := base.New(
		t,
		service_tester.WithInstances(
			constant.FixtureInstance,
			constant.FixtureOtherInstance,
		),
	).Served()
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
