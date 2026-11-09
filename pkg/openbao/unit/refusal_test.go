package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/openbao"
	"github.com/funtimecoding/soil/pkg/openbao/unit/upstream_tester"
	web "github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"net/http"
	"testing"
)

func TestSealedStoreIsUnavailable(t *testing.T) {
	s, _ := web.New(
		t,
		http.StatusServiceUnavailable,
		`{"errors":["Vault is sealed"]}`,
	)
	c := openbao.New(web.Locator(t, s.URL))
	defer c.Close()
	_, e := c.List("alfa")
	assert.True(t, unavailable.Is(e))
}

func TestRefusedPolicyIsCaptured(t *testing.T) {
	s, _ := web.New(t, http.StatusForbidden, `{"errors":["permission denied"]}`)
	c := openbao.New(web.Locator(t, s.URL))
	defer c.Close()
	_, e := c.List("alfa")
	assert.True(t, classify.Reportable(e))
}

func TestSealedStoreFailsProbe(t *testing.T) {
	s, _ := web.New(
		t,
		http.StatusOK,
		`{"initialized":true,"sealed":true,"version":"2.7.1"}`,
	)
	c := openbao.New(web.Locator(t, s.URL))
	defer c.Close()
	assert.True(t, unavailable.Is(c.Reachable(t.Context())))
}

func TestUninitializedStoreFailsProbe(t *testing.T) {
	s, _ := web.New(
		t,
		http.StatusOK,
		`{"initialized":false,"sealed":false,"version":"2.7.1"}`,
	)
	c := openbao.New(web.Locator(t, s.URL))
	defer c.Close()
	assert.True(t, unavailable.Is(c.Reachable(t.Context())))
}

func TestDeadLoginFailsProbe(t *testing.T) {
	assert.NotNil(
		t,
		client(t, upstream_tester.New(t), "bravo-jwt").Reachable(t.Context()),
	)
}
