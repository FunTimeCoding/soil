package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/integration/model_context_tester"
	"testing"
)

func TestEveryToolIsRegistered(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	assert.Count(t, 12, o.Client.ListTools())
}

func TestChainIsBuiltThroughTools(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.CreateRoot()
	assert.StringContains(t, "Example Issuing CA", o.CreateCluster())
	assert.StringContains(
		t,
		"Example",
		o.Client.MustCallTool(constant.ListAuthorities, map[string]any{}),
	)
}

func TestSecondRootIsRefusedThroughTools(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.CreateRoot()
	assert.StringContains(
		t,
		"An authority of that name is already live",
		o.CreateRootRefusal(),
	)
}

func TestIssuedCertificateReturnsItsKeyOnce(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.CreateRoot()
	o.CreateCluster()
	issued := o.Client.MustCallTool(
		constant.IssueCertificate,
		map[string]any{
			constant.AuthorityParameter:  constant.FixtureClusterAuthority,
			constant.KindParameter:       string(constant.KindServer),
			constant.CommonNameParameter: constant.FixtureCommonName,
			constant.HostParameter:       []any{constant.FixtureHost},
		},
	)
	assert.StringContains(t, "PRIVATE KEY", issued)
	assert.StringNotContains(
		t,
		"PRIVATE KEY",
		o.Client.MustCallTool(constant.ListCertificates, map[string]any{}),
	)
}

func TestRootCertificateToolServesTheAnchor(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.CreateRoot()
	assert.StringContains(
		t,
		"BEGIN CERTIFICATE",
		o.Client.MustCallTool(constant.RootCertificate, map[string]any{}),
	)
}

func TestPublishThroughToolsCommitsTheChain(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.CreateRoot()
	o.CreateCluster()
	o.Client.MustCallTool(constant.Publish, map[string]any{})
	assert.Integer(t, 1, len(o.Server.Forge.Commits()))
	assert.Integer(t, 6, len(o.Server.Forge.Commits()[0].Actions))
	assert.StringContains(
		t,
		"Nothing to publish",
		o.Client.MustCallTool(constant.Publish, map[string]any{}),
	)
}
