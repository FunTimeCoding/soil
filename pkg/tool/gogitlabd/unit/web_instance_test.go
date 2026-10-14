package unit

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/service_tester"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"testing"
)

func TestWebRefusesAnUnresolvableInstanceAsBadRequest(t *testing.T) {
	v := startServer(
		t,
		service_tester.Service(
			map[string]*mock_client.Client{
				constant.FixtureInstance:      mock_client.New(),
				constant.FixtureOtherInstance: mock_client.New(),
			},
		),
	)
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
