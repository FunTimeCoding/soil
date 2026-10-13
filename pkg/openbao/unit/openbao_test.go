package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/openbao/unit/upstream_tester"
	"testing"
)

func TestHealthNeedsNoLogin(t *testing.T) {
	u := upstream_tester.New(t)
	s := client(t, u, "alfa-jwt").MustHealth()
	assert.True(t, s.Initialized)
	assert.False(t, s.Sealed)
	assert.String(t, "2.7.1", s.Version)
	assert.Integer(t, 0, u.Logins())
}

func TestLoginHappensOnceAndCarriesTheToken(t *testing.T) {
	u := upstream_tester.New(t)
	c := client(t, u, "alfa-jwt")
	defer c.Close()
	c.MustList("alfa")
	c.MustRead("alfa/bravo")
	assert.Integer(t, 1, u.Logins())
	assert.Strings(t, []string{"", "alfa-token", "alfa-token"}, u.Tokens())
}

func TestRefusedLoginFailsTheCall(t *testing.T) {
	_, e := client(t, upstream_tester.New(t), "bravo-jwt").List("alfa")
	assert.StringContains(t, "log in:", e.Error())
}

func TestListSortsKeys(t *testing.T) {
	assert.Strings(
		t,
		[]string{"bravo", "charlie/"},
		client(t, upstream_tester.New(t), "alfa-jwt").MustList("alfa"),
	)
}

func TestListOfAMissingPathIsEmpty(t *testing.T) {
	assert.Count(
		t,
		0,
		client(t, upstream_tester.New(t), "alfa-jwt").MustList("missing"),
	)
}

func TestReadReturnsEveryKey(t *testing.T) {
	assert.Any(
		t,
		map[string]string{"DELTA": "echo", "FOXTROT": "golf"},
		client(t, upstream_tester.New(t), "alfa-jwt").MustRead("alfa/bravo"),
	)
}

func TestReadOfAMissingSecretIsNotFound(t *testing.T) {
	_, e := client(t, upstream_tester.New(t), "alfa-jwt").Read("alfa/missing")
	assert.True(t, not_found.Is(e))
}

func TestPutSendsTheValuesAsTheSecretBody(t *testing.T) {
	u := upstream_tester.New(t)
	client(t, u, "alfa-jwt").MustPut(
		"alfa/bravo",
		map[string]string{"DELTA": "echo"},
	)
	assert.Any(t, map[string]any{"DELTA": "echo"}, u.Written()["data"])
}
