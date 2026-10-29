package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/provision/salt/salt_tester"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTheFirstCallLogsInAndSendsTheToken(t *testing.T) {
	s := salt_tester.New(t, "bravo")
	c := newSaltClient(t, s.URL, "bravo")
	defer c.Close()
	assert.Integer(t, 0, s.Logins())
	keys, e := c.ListKeys()
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"alfa"}, keys.Minions)
	assert.Integer(t, 1, s.Logins())
	assert.Strings(t, []string{"token-1"}, s.Seen())
}

func TestExpiredTokenRenewsOnce(t *testing.T) {
	s := salt_tester.New(t, "bravo")
	c := newSaltClient(t, s.URL, "bravo")
	defer c.Close()
	_, e := c.ListKeys()
	assert.FatalOnError(t, e)
	s.Expire()
	_, f := c.ListKeys()
	assert.FatalOnError(t, f)
	assert.Integer(t, 2, s.Logins())
	assert.Strings(t, []string{"token-1", "token-1", "token-2"}, s.Seen())
}

func TestRefusalCarriesTheParagraph(t *testing.T) {
	s := salt_tester.New(t, "bravo")
	_, e := newSaltClient(t, s.URL, "bravo").ListJobs()
	assert.True(t, unexpected.Is(e))
	assert.StringContains(
		t,
		"salt-api status: 500: An unexpected error occurred",
		e.Error(),
	)
}

func TestAWrongPasswordIsRefusedWithTheReason(t *testing.T) {
	s := salt_tester.New(t, "bravo")
	c := newSaltClient(t, s.URL, "charlie")
	defer c.Close()
	e := c.Login(t.Context())
	assert.Nil(t, connection.Classify(e))
	assert.StringContains(
		t,
		"salt-api status: 401: Could not authenticate using provided credentials",
		e.Error(),
	)
	_, f := c.ListKeys()
	assert.StringContains(t, "Could not authenticate", f.Error())
}

func TestALoginNobodyAnswersIsAConnectionFailure(t *testing.T) {
	s := salt_tester.New(t, "bravo")
	c := newSaltClient(t, s.URL, "bravo")
	defer c.Close()
	s.Close()
	assert.NotNil(t, connection.Classify(c.Login(t.Context())))
}

func TestALoginThatHangsEndsWithItsContext(t *testing.T) {
	release := make(chan struct{})
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				http.ResponseWriter,
				*http.Request,
			) {
				<-release
			},
		),
	)
	defer s.Close()
	defer close(release)
	c := newSaltClient(t, s.URL, "bravo")
	defer c.Close()
	x, cancel := context.WithTimeout(t.Context(), 50*time.Millisecond)
	defer cancel()
	start := time.Now()
	e := c.Login(x)
	assert.True(t, time.Since(start) < time.Second)
	assert.NotNil(t, connection.Classify(e))
}
