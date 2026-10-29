package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/habitica"
	"github.com/funtimecoding/soil/pkg/habitica/face"
	"github.com/funtimecoding/soil/pkg/habitica/mock_client"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"net/http"
	"testing"
)

func TestTheHabiticaClientAndItsMockShareTheFace(t *testing.T) {
	c := habitica.New("habitica.example", "alfa-user", "bravo-token")
	defer c.Close()
	m := mock_client.New()
	defer m.Close()
	faces := []face.Client{c, m}
	assert.Count(t, 2, faces)
}

func TestARefusedUserReadReturnsTheError(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusUnauthorized,
		`{"success":false,"error":"NotAuthorized","message":"token expired"}`,
	)
	c := habitica.FromBasic(newClient(t, s.URL))
	defer c.Close()
	g, e := c.UserGear()
	assert.True(t, forbidden.Is(e))
	assert.True(t, g == nil)
	v, f := c.UserStats()
	assert.True(t, forbidden.Is(f))
	assert.True(t, v == nil)
	assert.True(t, forbidden.Is(c.Reachable(t.Context())))
}
