package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"testing"
)

func TestNewLogsInAndSendsTheToken(t *testing.T) {
	s := newSaltServer(t, "bravo")
	keys, e := newSaltClient(t, s.URL, "bravo").ListKeys()
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"alfa"}, keys.Minions)
	assert.Integer(t, 1, s.logins)
	assert.Strings(t, []string{"token-1"}, s.seen)
}

func TestExpiredTokenRenewsOnce(t *testing.T) {
	s := newSaltServer(t, "bravo")
	c := newSaltClient(t, s.URL, "bravo")
	s.expire()
	_, e := c.ListKeys()
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, s.logins)
	assert.Strings(t, []string{"token-1", "token-2"}, s.seen)
}

func TestRefusalCarriesTheParagraph(t *testing.T) {
	s := newSaltServer(t, "bravo")
	_, e := newSaltClient(t, s.URL, "bravo").ListJobs()
	assert.True(t, unexpected.Is(e))
	assert.StringContains(
		t,
		"salt-api status: 500: An unexpected error occurred",
		e.Error(),
	)
}

func TestWrongPasswordPanicsWithTheReason(t *testing.T) {
	s := newSaltServer(t, "bravo")
	defer func() {
		r := recover()
		e, okay := r.(error)
		assert.True(t, okay)
		assert.StringContains(
			t,
			"salt-api status: 401: Could not authenticate using provided credentials",
			e.Error(),
		)
	}()
	newSaltClient(t, s.URL, "charlie")
}
