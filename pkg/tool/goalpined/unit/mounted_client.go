package unit

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/client"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/unit/service_tester"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func mountedClient(
	t *testing.T,
	o *service_tester.Tester,
) *client.Client {
	t.Helper()
	v := mountedUpload(t, o.Service)
	result := client.New(
		locator.New(webConstant.Localhost).Insecure().Port(v.Port),
		constant.ModelContextTestToken,
	)
	t.Cleanup(result.Close)

	return result
}
