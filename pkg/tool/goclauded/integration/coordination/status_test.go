package coordination

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"io"
	"net/http"
	"testing"
)

func TestStatusReportsNothingWhenSound(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "status surface")
	assert.StringContains(
		t,
		"Nothing inconsistent.",
		a.MustCallTool(constant.Status, map[string]any{}),
	)
	c := s.RESTClient(t)
	response, e := c.GetStatusWithResponse(context.Background())
	assert.FatalOnError(t, e)
	assert.Count(t, 0, response.JSON200.Findings)
}

func TestStatusReportsFindingsAcrossSurfaces(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	assert.FatalOnError(
		t,
		s.Store.Store.PushQueue(
			"",
			"Nobody",
			constant.QueueTimeout,
			"orphaned",
		),
	)
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "status surface")
	result := a.MustCallTool(constant.Status, map[string]any{})
	assert.StringContains(t, "unowned_queue", result)
	assert.StringContains(t, "Nobody", result)
	c := s.RESTClient(t)
	response, e := c.GetStatusWithResponse(context.Background())
	assert.FatalOnError(t, e)
	var unowned []client.FindingEntry

	for _, f := range response.JSON200.Findings {
		if f.Kind == constant.UnownedQueue {
			unowned = append(unowned, f)
		}
	}

	assert.Count(t, 1, unowned)
	assert.StringContains(t, "Nobody", unowned[0].Detail)
	page, f := http.Get(
		fmt.Sprintf("http://localhost:%d%s", s.Port, constant.StatusPath),
	)
	assert.FatalOnError(t, f)

	defer errors.PanicClose(page.Body)
	body, g := io.ReadAll(page.Body)
	assert.FatalOnError(t, g)
	assert.StringContains(t, "unowned_queue", string(body))
	assert.StringContains(t, "Nobody", string(body))
}
