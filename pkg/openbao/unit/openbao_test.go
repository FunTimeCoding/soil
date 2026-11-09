package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/openbao/unit/upstream_tester"
	"testing"
)

func TestHealthNeedsNoLogin(t *testing.T) {
	u := upstream_tester.New(t)
	s, e := client(t, u, "alfa-jwt").Health()
	assert.FatalOnError(t, e)
	assert.True(t, s.Initialized)
	assert.False(t, s.Sealed)
	assert.String(t, "2.7.1", s.Version)
	assert.Integer(t, 0, u.Logins())
}

func TestLoginHappensOnceAndCarriesToken(t *testing.T) {
	u := upstream_tester.New(t)
	c := client(t, u, "alfa-jwt")
	defer c.Close()
	_, e := c.List("alfa")
	assert.FatalOnError(t, e)
	_, f := c.Read("alfa/bravo")
	assert.FatalOnError(t, f)
	assert.Integer(t, 1, u.Logins())
	assert.Strings(t, []string{"", "alfa-token", "alfa-token"}, u.Tokens())
}

func TestRefusedLoginFailsCall(t *testing.T) {
	_, e := client(t, upstream_tester.New(t), "bravo-jwt").List("alfa")
	assert.StringContains(t, "log in:", e.Error())
}

func TestListSortsKeys(t *testing.T) {
	keys, e := client(t, upstream_tester.New(t), "alfa-jwt").List("alfa")
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"bravo", "charlie/"}, keys)
}

func TestListOfMissingPathIsEmpty(t *testing.T) {
	keys, e := client(t, upstream_tester.New(t), "alfa-jwt").List("missing")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, keys)
}

func TestReadReturnsEveryKey(t *testing.T) {
	values, e := client(t, upstream_tester.New(t), "alfa-jwt").Read(
		"alfa/bravo",
	)
	assert.FatalOnError(t, e)
	assert.Any(t, map[string]string{"DELTA": "echo", "FOXTROT": "golf"}, values)
}

func TestReadOfMissingSecretIsNotFound(t *testing.T) {
	_, e := client(t, upstream_tester.New(t), "alfa-jwt").Read("alfa/missing")
	assert.True(t, not_found.Is(e))
}

func TestPutSendsValuesAsSecretBody(t *testing.T) {
	u := upstream_tester.New(t)
	assert.FatalOnError(
		t,
		client(t, u, "alfa-jwt").Put(
			"alfa/bravo",
			map[string]string{"DELTA": "echo"},
		),
	)
	assert.Any(t, map[string]any{"DELTA": "echo"}, u.Written()["data"])
}
