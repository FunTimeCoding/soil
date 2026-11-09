package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	soilErrors "github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/dropped"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	gomonitord "github.com/funtimecoding/soil/pkg/tool/gomonitord/client"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"
	"testing"
)

func TestMockClaimsReleasesAndStreams(t *testing.T) {
	m := mock_client.New()
	defer m.Close()
	x := t.Context()
	_, e := m.PutClaim(x, "jira-ABC-1", "alfa@host.example")
	assert.FatalOnError(t, e)
	_, f := m.PutClaim(x, "jira-ABC-1", "alfa@host.example")
	assert.FatalOnError(t, f)
	claims, g := m.GetClaims(x)
	assert.FatalOnError(t, g)
	assert.Count(t, 1, claims)
	b, h := m.Listen(x)
	assert.FatalOnError(t, h)
	defer soilErrors.PanicClose(b)
	var streamed []client.Claim
	i := gomonitord.ReadClaims(b, func(v []client.Claim) { streamed = v })
	assert.True(t, dropped.Is(i))
	assert.Count(t, 1, streamed)
	assert.FatalOnError(t, m.DeleteClaim(x, "jira-ABC-1", "alfa@host.example"))
	after, j := m.GetClaims(x)
	assert.FatalOnError(t, j)
	assert.Count(t, 0, after)
	_, k := m.PutClaim(x, "jira-ABC-1", "")
	assert.True(t, validation.Is(k))
}

func TestMockFailsEveryOperation(t *testing.T) {
	m := mock_client.New()
	defer m.Close()
	m.Failure = errors.New("alfa")
	x := t.Context()
	assert.NotNil(t, m.Reachable(x))
	_, e := m.GetClaims(x)
	assert.NotNil(t, e)
	_, f := m.PutClaim(x, "jira-ABC-1", "alfa@host.example")
	assert.NotNil(t, f)
	assert.NotNil(t, m.DeleteClaim(x, "jira-ABC-1", "alfa@host.example"))
	_, g := m.Listen(x)
	assert.NotNil(t, g)
}
