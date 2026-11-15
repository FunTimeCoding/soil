package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/openbao"
	"github.com/funtimecoding/soil/pkg/openbao/unit/openbao_server"
	"github.com/funtimecoding/soil/pkg/openbao/unit/openbao_tester"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func TestSealedStoreIsUnavailable(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusServiceUnavailable),
		upstream_server.WithAnswer(`{"errors":["Vault is sealed"]}`),
	)
	c := openbao.New(locator.NewOrigin(s.URL))
	defer c.Close()
	_, e := c.List("alfa")
	assert.True(t, unavailable.Is(e))
}

func TestRefusedPolicyIsCaptured(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusForbidden),
		upstream_server.WithAnswer(`{"errors":["permission denied"]}`),
	)
	c := openbao.New(locator.NewOrigin(s.URL))
	defer c.Close()
	_, e := c.List("alfa")
	assert.True(t, classify.Reportable(e))
}

func TestSealedStoreFailsProbe(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(
			`{"initialized":true,"sealed":true,"version":"2.7.1"}`,
		),
	)
	c := openbao.New(locator.NewOrigin(s.URL))
	defer c.Close()
	assert.True(t, unavailable.Is(c.Reachable(t.Context())))
}

func TestUninitializedStoreFailsProbe(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(
			`{"initialized":false,"sealed":false,"version":"2.7.1"}`,
		),
	)
	c := openbao.New(locator.NewOrigin(s.URL))
	defer c.Close()
	assert.True(t, unavailable.Is(c.Reachable(t.Context())))
}

func TestDeadLoginFailsProbe(t *testing.T) {
	assert.NotNil(
		t,
		openbao_tester.New(t, openbao_server.New(t), "bravo-jwt").Reachable(
			t.Context(),
		),
	)
}
