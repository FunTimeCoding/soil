package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/openbao/unit/openbao_server"
	"github.com/funtimecoding/soil/pkg/openbao/unit/openbao_tester"
	"testing"
)

func TestHealthNeedsNoLogin(t *testing.T) {
	r := openbao_server.New(t)
	s, e := openbao_tester.New(t, r, "alfa-jwt").Health()
	assert.FatalOnError(t, e)
	assert.True(t, s.Initialized)
	assert.False(t, s.Sealed)
	assert.String(t, "2.7.1", s.Version)
	assert.Integer(t, 0, r.Logins())
}

func TestLoginHappensOnceAndCarriesToken(t *testing.T) {
	s := openbao_server.New(t)
	c := openbao_tester.New(t, s, "alfa-jwt")
	defer c.Close()
	_, e := c.List("alfa")
	assert.FatalOnError(t, e)
	_, f := c.Read("alfa/bravo")
	assert.FatalOnError(t, f)
	assert.Integer(t, 1, s.Logins())
	assert.Strings(t, []string{"", "alfa-token", "alfa-token"}, s.Tokens())
}

func TestRefusedLoginFailsCall(t *testing.T) {
	_, e := openbao_tester.New(t, openbao_server.New(t), "bravo-jwt").List(
		"alfa",
	)
	assert.StringContains(t, "log in:", e.Error())
}

func TestListSortsKeys(t *testing.T) {
	keys, e := openbao_tester.New(t, openbao_server.New(t), "alfa-jwt").List(
		"alfa",
	)
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"bravo", "charlie/"}, keys)
}

func TestListOfMissingPathIsEmpty(t *testing.T) {
	keys, e := openbao_tester.New(t, openbao_server.New(t), "alfa-jwt").List(
		"missing",
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, keys)
}

func TestReadReturnsEveryKey(t *testing.T) {
	values, e := openbao_tester.New(t, openbao_server.New(t), "alfa-jwt").Read(
		"alfa/bravo",
	)
	assert.FatalOnError(t, e)
	assert.Any(t, map[string]string{"DELTA": "echo", "FOXTROT": "golf"}, values)
}

func TestReadOfMissingSecretIsNotFound(t *testing.T) {
	_, e := openbao_tester.New(t, openbao_server.New(t), "alfa-jwt").Read(
		"alfa/missing",
	)
	assert.True(t, not_found.Is(e))
}

func TestPutSendsValuesAsSecretBody(t *testing.T) {
	s := openbao_server.New(t)
	assert.FatalOnError(
		t,
		openbao_tester.New(t, s, "alfa-jwt").Put(
			"alfa/bravo",
			map[string]string{"DELTA": "echo"},
		),
	)
	assert.Any(t, map[string]any{"DELTA": "echo"}, s.Written()["data"])
}
