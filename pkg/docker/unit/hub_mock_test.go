package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/docker/face"
	"github.com/funtimecoding/soil/pkg/docker/hub"
	"github.com/funtimecoding/soil/pkg/docker/hub/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"testing"
)

func TestHubClientsSatisfyFace(t *testing.T) {
	var c face.DockerHubClient = hub.New()
	defer c.Close()
	var m face.DockerHubClient = mock_client.New()
	defer m.Close()
	assert.NotNil(t, m)
}

func TestMockAnswersSeededTagsAndMissesLikeClient(t *testing.T) {
	m := mock_client.New()
	defer m.Close()
	m.SeedTags("library/foxtrot", "1.0-alpine")
	tags, e := m.Tags("library/foxtrot")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, tags)
	assert.String(t, "1.0-alpine", tags[0].Name)
	_, e = m.Tags("library/golf")
	assert.True(t, not_found.Is(e))
}
