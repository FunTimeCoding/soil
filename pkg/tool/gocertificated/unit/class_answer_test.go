package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/server_tester"
	"net/http"
	"testing"
)

func TestMissingAuthorityAnswersNotFoundQuietly(t *testing.T) {
	s := base.New(t)
	r, e := s.GeneratedClient().GetAuthorityWithResponse(
		t.Context(),
		constant.FixtureClusterAuthority,
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, r.StatusCode())
	assert.FatalNotNil(t, r.JSON404)
	assert.StringContains(t, "not found", r.JSON404.Error)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestRootBeforeCreationAnswersNotFoundQuietly(t *testing.T) {
	s := base.New(t)
	r, e := s.GeneratedClient().GetRootWithResponse(t.Context())
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, r.StatusCode())
	assert.NotNil(t, r.JSON404)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestRootFaultIsCapturedAsServerError(t *testing.T) {
	s := base.New(t)
	c := s.GeneratedClient()
	s.Daemon().Close()
	r, e := c.GetRootWithResponse(t.Context())
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusInternalServerError, r.StatusCode())
	assert.Count(t, 1, s.Reporter.Events())
}

func TestSecondRootAnswersConflictQuietly(t *testing.T) {
	s := base.New(t)
	c := s.GeneratedClient()
	server_tester.PostAuthority(t, c, server_tester.RootBody())
	r := server_tester.PostAuthority(t, c, server_tester.RootBody())
	assert.Integer(t, http.StatusConflict, r.StatusCode())
	assert.FatalNotNil(t, r.JSON409)
	assert.StringContains(t, "already live", r.JSON409.Error)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestUnusableAddressAnswersBadRequestQuietly(t *testing.T) {
	s := base.New(t)
	c := s.GeneratedClient()
	server_tester.PostAuthority(t, c, server_tester.RootBody())
	b := server_tester.ClusterBody()
	b.PermittedAddress = &[]string{"nowhere"}
	r := server_tester.PostAuthority(t, c, b)
	assert.Integer(t, http.StatusBadRequest, r.StatusCode())
	assert.FatalNotNil(t, r.JSON400)
	assert.StringContains(t, "nowhere", r.JSON400.Error)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestIssuingFromMissingAuthorityAnswersNotFoundQuietly(t *testing.T) {
	s := base.New(t)
	r := server_tester.PostCertificate(t, s.GeneratedClient())
	assert.Integer(t, http.StatusNotFound, r.StatusCode())
	assert.NotNil(t, r.JSON404)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestRevokingTwiceAnswersNotFoundQuietly(t *testing.T) {
	s := base.New(t)
	c := s.GeneratedClient()
	server_tester.PostAuthority(t, c, server_tester.RootBody())
	server_tester.PostAuthority(t, c, server_tester.ClusterBody())
	issued := server_tester.PostCertificate(t, c)
	assert.FatalNotNil(t, issued.JSON200)
	serial := issued.JSON200.Serial
	_, e := c.PostRevocationWithResponse(t.Context(), serial)
	assert.FatalOnError(t, e)
	r, f := c.PostRevocationWithResponse(t.Context(), serial)
	assert.FatalOnError(t, f)
	assert.Integer(t, http.StatusNotFound, r.StatusCode())
	assert.Count(t, 0, s.Reporter.Events())
}

func TestMissingCertificateRefusesThroughToolsQuietly(t *testing.T) {
	s := base.New(t)
	assert.StringContains(
		t,
		"not found",
		s.Tools().MustCallToolError(
			constant.GetCertificate,
			map[string]any{constant.SerialParameter: "absent"},
		),
	)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestToolFaultIsCaptured(t *testing.T) {
	s := base.New(t)
	c := s.Tools()
	s.Daemon().Close()
	c.MustCallToolError(constant.RootCertificate, map[string]any{})
	assert.Count(t, 1, s.Reporter.Events())
}

func TestRootPathAnswersMissingRootWithNotFound(t *testing.T) {
	s := base.New(t)
	s.Served().VerifyStatus(t, constant.RootPath, http.StatusNotFound)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestFormRejectionIsNotCaptured(t *testing.T) {
	s := base.New(t)
	s.Submit(constant.CreateAuthorityPath, base.RootValues())
	s.Submit(constant.CreateAuthorityPath, base.RootValues())
	assert.Count(t, 0, s.Reporter.Events())
}

func TestFormFaultIsCapturedNotInlined(t *testing.T) {
	s := base.New(t)
	s.Served()
	s.Daemon().Close()
	page := s.Submit(constant.CreateAuthorityPath, base.RootValues())
	assert.Count(t, 1, s.Reporter.Events())
	assert.StringNotContains(t, "Example Root CA", page)
}
