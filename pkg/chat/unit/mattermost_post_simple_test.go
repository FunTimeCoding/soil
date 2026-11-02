package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/mattermost/mattermost/server/public/model"
	"testing"
)

func TestTheMockListsAndDeletesAChannelsPosts(t *testing.T) {
	m := mock_client.New("charlie")
	defer m.Close()
	m.AddChannelPost("alfa", "bravo")
	m.AddChannelPost("delta", "echo")
	posts, e := m.Posts(m.MustTeamChannel("alfa"))
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"bravo"}, posts.Order)
	assert.FatalOnError(t, m.DeletePost(posts.Posts["bravo"]))
	assert.True(t, not_found.Is(m.DeletePost(posts.Posts["bravo"])))
	_, f := m.FindPost("echo")
	assert.FatalOnError(t, f)
}

func TestTheMockPostsToALookedUpChannelByIdentifier(t *testing.T) {
	m := mock_client.New("charlie")
	defer m.Close()

	for _, h := range []*model.Channel{
		m.MustChannelByName(nil, "alfa"),
		m.MustTeamChannel("alfa"),
		m.MustChannel("alfa"),
	} {
		m.MustPostSimple(h, "bravo")
	}

	assert.Count(t, 3, m.Posted)

	for _, p := range m.Posted {
		assert.String(t, "alfa", p.ChannelId)
		assert.String(t, "bravo", p.Message)
	}
}
