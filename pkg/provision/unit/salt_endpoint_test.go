package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/provision/salt/salt_tester"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFirstCallLogsInAndSendsToken(t *testing.T) {
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

func TestRefusalCarriesParagraph(t *testing.T) {
	s := salt_tester.New(t, "bravo")
	_, e := newSaltClient(t, s.URL, "bravo").ListJobs()
	assert.True(t, unexpected.Is(e))
	assert.StringContains(
		t,
		"salt-api status: 500: An unexpected error occurred",
		e.Error(),
	)
}

func TestGatewayTimeoutIsUnavailable(t *testing.T) {
	s := salt_tester.New(t, "bravo")
	c := newSaltClient(t, s.URL, "bravo")
	assert.FatalOnError(t, c.Login(t.Context()))
	s.Gateway(http.StatusGatewayTimeout)
	_, e := c.ListJobs()
	assert.True(t, unavailable.Is(e))
	assert.StringContains(t, "salt-api status: 504", e.Error())
}

func TestBadGatewayIsUnavailable(t *testing.T) {
	s := salt_tester.New(t, "bravo")
	c := newSaltClient(t, s.URL, "bravo")
	assert.FatalOnError(t, c.Login(t.Context()))
	s.Gateway(http.StatusBadGateway)
	_, e := c.ListJobs()
	assert.True(t, unavailable.Is(e))
}

func TestClassifiedStatusWinsOverParagraph(t *testing.T) {
	s := salt_tester.New(t, "bravo")
	c := newSaltClient(t, s.URL, "bravo")
	assert.FatalOnError(t, c.Login(t.Context()))
	s.Refuse(http.StatusServiceUnavailable, "salt master is restarting")
	_, e := c.ListJobs()
	assert.True(t, unavailable.Is(e))
	assert.StringContains(
		t,
		"salt-api status: 503: salt master is restarting",
		e.Error(),
	)
}

func TestWrongPasswordIsRefusedWithReason(t *testing.T) {
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

func TestLoginNobodyAnswersIsConnectionFailure(t *testing.T) {
	s := salt_tester.New(t, "bravo")
	c := newSaltClient(t, s.URL, "bravo")
	defer c.Close()
	s.Close()
	assert.NotNil(t, connection.Classify(c.Login(t.Context())))
}

func TestPostThatHangsEndsWithItsContext(t *testing.T) {
	s := salt_tester.New(t, "bravo")
	c := newSaltClient(t, s.URL, "bravo")
	defer c.Close()
	assert.FatalOnError(t, c.Login(t.Context()))
	s.HoldPost()
	defer s.Release()
	x, cancel := context.WithCancel(t.Context())
	posted := make(chan error, 1)
	go func() {
		var out any
		posted <- c.Post(x, "", map[string]string{}, &out)
	}()

	select {
	case <-s.Holding():
	case <-time.After(10 * time.Second):
		t.Fatal("the post never reached the master")
	}

	cancel()

	select {
	case e := <-posted:
		assert.NotNil(t, e)
	case <-time.After(10 * time.Second):
		t.Fatal("the post did not end with its context")
	}
}

func TestLoginThatHangsEndsWithItsContext(t *testing.T) {
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
