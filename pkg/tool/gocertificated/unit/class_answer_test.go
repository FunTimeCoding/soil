package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/model_context_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/server_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/web_interface_tester"
	"net/http"
	"testing"
)

func TestMissingAuthorityAnswersNotFoundQuietly(t *testing.T) {
	s := base.New(t)
	r, e := restClient(t, s).GetAuthorityWithResponse(
		t.Context(),
		constant.FixtureClusterAuthority,
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, r.StatusCode())
	assert.StringContains(t, "not found", r.JSON404.Error)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestRootBeforeCreationAnswersNotFoundQuietly(t *testing.T) {
	s := base.New(t)
	r, e := restClient(t, s).GetRootWithResponse(t.Context())
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, r.StatusCode())
	assert.Count(t, 0, s.Reporter.Events())
}

func TestRootFaultIsCapturedAsServerError(t *testing.T) {
	s := base.New(t)
	s.Service.Close()
	r, e := restClient(t, s).GetRootWithResponse(t.Context())
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusInternalServerError, r.StatusCode())
	assert.Count(t, 1, s.Reporter.Events())
}

func TestSecondRootAnswersConflictQuietly(t *testing.T) {
	s := base.New(t)
	c := restClient(t, s)
	postAuthority(t, c, server_tester.RootBody())
	r := postAuthority(t, c, server_tester.RootBody())
	assert.Integer(t, http.StatusConflict, r.StatusCode())
	assert.StringContains(t, "already live", r.JSON409.Error)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestUnusableAddressAnswersBadRequestQuietly(t *testing.T) {
	s := base.New(t)
	c := restClient(t, s)
	postAuthority(t, c, server_tester.RootBody())
	b := server_tester.ClusterBody()
	b.PermittedAddress = &[]string{"nowhere"}
	r := postAuthority(t, c, b)
	assert.Integer(t, http.StatusBadRequest, r.StatusCode())
	assert.StringContains(t, "nowhere", r.JSON400.Error)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestIssuingFromMissingAuthorityAnswersNotFoundQuietly(t *testing.T) {
	s := base.New(t)
	r := postCertificate(t, restClient(t, s))
	assert.Integer(t, http.StatusNotFound, r.StatusCode())
	assert.Count(t, 0, s.Reporter.Events())
}

func TestRevokingTwiceAnswersNotFoundQuietly(t *testing.T) {
	s := base.New(t)
	c := restClient(t, s)
	postAuthority(t, c, server_tester.RootBody())
	postAuthority(t, c, server_tester.ClusterBody())
	serial := postCertificate(t, c).JSON200.Serial
	_, e := c.PostRevocationWithResponse(t.Context(), serial)
	assert.FatalOnError(t, e)
	r, f := c.PostRevocationWithResponse(t.Context(), serial)
	assert.FatalOnError(t, f)
	assert.Integer(t, http.StatusNotFound, r.StatusCode())
	assert.Count(t, 0, s.Reporter.Events())
}

func TestMissingCertificateRefusesThroughToolsQuietly(t *testing.T) {
	o := model_context_tester.New(t)
	assert.StringContains(
		t,
		"not found",
		o.Client.MustCallToolError(
			constant.GetCertificate,
			map[string]any{constant.SerialParameter: "absent"},
		),
	)
	assert.Count(t, 0, o.Server.Reporter.Events())
}

func TestToolFaultIsCaptured(t *testing.T) {
	o := model_context_tester.New(t)
	o.Server.Service.Close()
	o.Client.MustCallToolError(constant.RootCertificate, map[string]any{})
	assert.Count(t, 1, o.Server.Reporter.Events())
}

func TestRootPathAnswersMissingRootWithNotFound(t *testing.T) {
	s := base.New(t)
	s.VerifyStatus(t, constant.RootPath, http.StatusNotFound)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestFormRejectionIsNotCaptured(t *testing.T) {
	o := web_interface_tester.New(t)
	o.Submit(constant.CreateAuthorityPath, rootValues())
	o.Submit(constant.CreateAuthorityPath, rootValues())
	assert.Count(t, 0, o.Server.Reporter.Events())
}

func TestFormFaultIsCapturedNotInlined(t *testing.T) {
	o := web_interface_tester.New(t)
	o.Server.Service.Close()
	page := o.Submit(constant.CreateAuthorityPath, rootValues())
	assert.Count(t, 1, o.Server.Reporter.Events())
	assert.StringNotContains(t, "Example Root CA", page)
}
