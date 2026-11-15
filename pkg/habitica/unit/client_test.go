package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/habitica"
	"github.com/funtimecoding/soil/pkg/habitica/face"
	"github.com/funtimecoding/soil/pkg/habitica/mock_client"
	"github.com/funtimecoding/soil/pkg/habitica/unit/habitica_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func TestHabiticaClientAndItsMockShareFace(t *testing.T) {
	c := habitica.New("habitica.example", "alfa-user", "bravo-token")
	defer c.Close()
	m := mock_client.New()
	defer m.Close()
	faces := []face.HabiticaClient{c, m}
	assert.Count(t, 2, faces)
}

func TestRefusedUserReadReturnsError(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusUnauthorized),
		upstream_server.WithAnswer(
			`{"success":false,"error":"NotAuthorized","message":"token expired"}`,
		),
	)
	c := habitica.FromBasic(habitica_tester.New(t, s.URL))
	defer c.Close()
	g, e := c.UserGear()
	assert.True(t, unexpected.Is(e))
	assert.True(t, g == nil)
	v, f := c.UserStats()
	assert.True(t, unexpected.Is(f))
	assert.True(t, v == nil)
	assert.True(t, unexpected.Is(c.Reachable(t.Context())))
}
