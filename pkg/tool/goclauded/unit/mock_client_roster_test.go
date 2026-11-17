package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"testing"
)

func TestMockEmptyRosterMatchesClient(t *testing.T) {
	client, e := base.New(t).Client().RecentSessions(25)
	assert.FatalOnError(t, e)
	assert.NotNil(t, client)
	assert.Count(t, 0, client)
	mock, f := mock_client.New(t).RecentSessions(25)
	assert.FatalOnError(t, f)
	assert.NotNil(t, mock)
	assert.Count(t, 0, mock)
}
