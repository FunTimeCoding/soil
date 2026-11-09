package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
	"testing"
)

func TestMockRecordsEachWrite(t *testing.T) {
	m := mock_client.New()
	defer m.Close()
	assert.Count(t, 0, writeEach(t, m))
	assert.Count(t, 1, m.Indexed)
	assert.String(t, "alfa", *m.Indexed[0].Collection)
	assert.Integer(t, 1, m.Embedded)
	assert.Integer(t, 1, m.Rechunked)
	assert.Count(t, 1, m.Tagged)
	assert.String(t, "bravo.md", m.Tagged[0].Path)
	assert.Count(t, 1, m.CollectionsAdded)
	assert.Count(t, 1, m.CollectionsDeleted)
	assert.String(t, "alfa", m.CollectionsDeleted[0].Name)
	assert.Count(t, 1, m.ContextsAdded)
	assert.Count(t, 1, m.ContextsDeleted)
	assert.String(t, "charlie", m.ContextsDeleted[0].PathPrefix)
	assert.Count(t, 0, m.Pushed)
	assert.Count(t, 0, m.Deleted)
}

func TestFailingMockRecordsNoWrite(t *testing.T) {
	m := mock_client.New()
	defer m.Close()
	m.Failure = errors.New("down")
	assert.Count(t, 8, writeEach(t, m))
	assert.Count(t, 0, m.Indexed)
	assert.Integer(t, 0, m.Embedded)
	assert.Integer(t, 0, m.Rechunked)
	assert.Count(t, 0, m.Tagged)
	assert.Count(t, 0, m.CollectionsAdded)
	assert.Count(t, 0, m.CollectionsDeleted)
	assert.Count(t, 0, m.ContextsAdded)
	assert.Count(t, 0, m.ContextsDeleted)
}
