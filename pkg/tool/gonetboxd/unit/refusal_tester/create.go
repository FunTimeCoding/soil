package refusal_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/refusing_source"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/service_tester"
	"testing"
)

func Create(
	t *testing.T,
	refusal error,
) (int, int) {
	t.Helper()
	s := base.New(t, service_tester.WithClient(refusing_source.New(t, refusal)))
	response, e := s.GeneratedClient().CreateSiteWithResponse(
		context.Background(),
		client.CreateSiteJSONRequestBody{Name: "alfa"},
	)
	assert.FatalOnError(t, e)

	return response.StatusCode(), len(s.Reporter.Events())
}
