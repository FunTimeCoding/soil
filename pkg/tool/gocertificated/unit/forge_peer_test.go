package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/server_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"net/http"
	"testing"
)

func TestForgeOutageAnswersUnavailableQuietly(t *testing.T) {
	s := base.New(t)
	c := s.GeneratedClient()
	server_tester.New(t).PostAuthority(c, server_tester.RootBody())
	s.Forge.Fail(outage_tester.RefusedDial())
	r, e := c.PostPublicationWithResponse(t.Context())
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusServiceUnavailable, r.StatusCode())
	assert.Count(t, 0, s.Reporter.Events())
}

func TestForgeOutageLeavesIssuingServing(t *testing.T) {
	s := base.New(t)
	c := s.GeneratedClient()
	server_tester.New(t).PostAuthority(c, server_tester.RootBody())
	server_tester.New(t).PostAuthority(c, server_tester.ClusterBody())
	s.Forge.Fail(outage_tester.RefusedDial())
	_, e := c.PostPublicationWithResponse(t.Context())
	assert.FatalOnError(t, e)
	r := server_tester.New(t).PostCertificate(c)
	assert.Integer(t, http.StatusOK, r.StatusCode())
}

func TestForgeFaultIsCaptured(t *testing.T) {
	s := base.New(t)
	c := s.GeneratedClient()
	server_tester.New(t).PostAuthority(c, server_tester.RootBody())
	s.Forge.Fail(errors.Format("commit rejected", constant.PublishBranch))
	r, e := c.PostPublicationWithResponse(t.Context())
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusInternalServerError, r.StatusCode())
	assert.Count(t, 1, s.Reporter.Events())
}

func TestDaemonStartsWhileForgeIsDown(t *testing.T) {
	s := base.New(t)
	s.Forge.Fail(outage_tester.RefusedDial())
	_, guard := s.Daemon().Workers()[0].(*peer.Guard)
	assert.True(t, guard)
}

func TestForgeOutageRefusesToolQuietly(t *testing.T) {
	s := base.New(t)
	s.CreateRoot()
	s.Forge.Fail(outage_tester.RefusedDial())
	s.Tools().MustCallToolError(constant.Publish, map[string]any{})
	assert.Count(t, 0, s.Reporter.Events())
}
