package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/client"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/face"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestTheClientReadsADiff(t *testing.T) {
	r := diffRepository(t)
	c := mountedClient(t, diffService(r.Clone), constant.ModelContextTestToken)
	assert.FatalOnError(t, c.Reachable(t.Context()))
	d, e := c.GetDiff(t.Context(), "echo", "", "")
	assert.FatalOnError(t, e)
	assert.Integer(t, 1, len(d.Entries))
	assert.String(t, "body", d.Entries[0].Class)
	assert.String(t, ". alfa Bravo", d.Entries[0].To.Key)
	assert.String(t, "alfa.go", d.Entries[0].To.File)
	assert.Integer(t, 3, d.Entries[0].To.Line)
	assert.String(t, "unchanged", d.Counts[0].Class)
	assert.Integer(t, 1, d.Counts[0].Count)
}

func TestTheClientHearsAnUnknownModule(t *testing.T) {
	r := diffRepository(t)
	c := mountedClient(t, diffService(r.Clone), constant.ModelContextTestToken)
	_, e := c.GetDiff(t.Context(), "golf", "", "")
	assert.True(t, not_found.Is(e))
}

func TestARefusedTokenIsAnError(t *testing.T) {
	r := diffRepository(t)
	c := mountedClient(t, diffService(r.Clone), "hotel-wrong")
	_, e := c.GetDiff(t.Context(), "echo", "", "")
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "diff status: 401", e.Error())
	assert.NotNil(t, c.Reachable(t.Context()))
}

func TestTheClientAndItsMockKeepTheFace(t *testing.T) {
	r := client.New(locator.New("localhost").Insecure(), "india")
	m := mock_client.New()
	clients := []face.Client{r, m}

	for _, c := range clients {
		c.Close()
	}

	assert.Integer(t, 1, m.Closed)
}
