package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/model_context_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/server_tester"
	"net/http"
	"testing"
)

func TestForgeOutageAnswersUnavailableQuietly(t *testing.T) {
	s := base.New(t)
	c := restClient(t, s)
	postAuthority(t, c, server_tester.RootBody())
	s.Forge.Fail(refusedConnection())
	r, e := c.PostPublicationWithResponse(t.Context())
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusServiceUnavailable, r.StatusCode())
	assert.Count(t, 0, s.Reporter.Events())
}

func TestForgeOutageLeavesIssuingServing(t *testing.T) {
	s := base.New(t)
	c := restClient(t, s)
	postAuthority(t, c, server_tester.RootBody())
	postAuthority(t, c, server_tester.ClusterBody())
	s.Forge.Fail(refusedConnection())
	_, e := c.PostPublicationWithResponse(t.Context())
	assert.FatalOnError(t, e)
	r := postCertificate(t, c)
	assert.Integer(t, http.StatusOK, r.StatusCode())
}

func TestForgeFaultIsCaptured(t *testing.T) {
	s := base.New(t)
	c := restClient(t, s)
	postAuthority(t, c, server_tester.RootBody())
	s.Forge.Fail(errors.Format("commit rejected", constant.PublishBranch))
	r, e := c.PostPublicationWithResponse(t.Context())
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusInternalServerError, r.StatusCode())
	assert.Count(t, 1, s.Reporter.Events())
}

func TestForgeOutageRefusesTheToolQuietly(t *testing.T) {
	o := model_context_tester.New(t)
	o.CreateRoot()
	o.Server.Forge.Fail(refusedConnection())
	o.Client.MustCallToolError(constant.Publish, map[string]any{})
	assert.Count(t, 0, o.Server.Reporter.Events())
}
