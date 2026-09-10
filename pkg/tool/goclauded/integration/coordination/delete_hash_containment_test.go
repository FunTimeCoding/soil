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

func TestDeleteHashRendersOnDetailPage(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	identifier := "11111111-2222-3333-4444-555555555555"
	s.Store.EnsureSession(identifier)
	s.Store.Store.UpdateFields(
		identifier,
		map[string]any{"turn_count": 4, "lines": 20},
	)
	hash := s.Service.DeleteHash(identifier)
	response, e := http.Get(
		fmt.Sprintf("http://localhost:%d/sessions/%s", s.Port, identifier),
	)
	assert.FatalOnError(t, e)

	defer errors.PanicClose(response.Body)
	body, f := io.ReadAll(response.Body)
	assert.FatalOnError(t, f)
	assert.StringContains(t, hash, string(body))
}

func TestDeleteHashStaysOutOfMachineSurfaces(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	s.Store.EnsureSession("11111111-2222-3333-4444-555555555555")
	s.Store.Store.UpdateFields(
		"11111111-2222-3333-4444-555555555555",
		map[string]any{"turn_count": 4, "lines": 20},
	)
	hash := s.Service.DeleteHash("11111111-2222-3333-4444-555555555555")
	assert.Integer(t, 8, len(hash))
	c := s.RESTClient(t)
	detail, e := c.GetSessionDetailWithResponse(
		context.Background(),
		"11111111-2222-3333-4444-555555555555",
	)
	assert.FatalOnError(t, e)
	assert.StringNotContains(t, hash, string(detail.Body))
	list, f := c.GetSessionsWithResponse(
		context.Background(),
		&client.GetSessionsParams{},
	)
	assert.FatalOnError(t, f)
	assert.StringNotContains(t, hash, string(list.Body))
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "hash containment")
	assert.StringNotContains(
		t,
		hash,
		a.MustCallTool(constant.ListSessions, map[string]any{}),
	)
	assert.StringNotContains(
		t,
		hash,
		a.MustCallTool(constant.SessionStatus, map[string]any{}),
	)
	assert.StringNotContains(
		t,
		hash,
		a.MustCallTool(constant.Roster, map[string]any{}),
	)
}
