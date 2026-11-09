package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/mattermost/mattermost/server/public/model"
	"testing"
)

func TestMockListsAndDeletesChannelsPosts(t *testing.T) {
	m := mock_client.New("charlie")
	defer m.Close()
	m.AddChannelPost("alfa", "bravo")
	m.AddChannelPost("delta", "echo")
	posts, e := m.RecentPosts(m.MustTeamChannel("alfa"), 0)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, posts)
	assert.String(t, "bravo", posts[0].Raw.Id)
	assert.FatalOnError(t, m.DeletePost(posts[0].Raw))
	assert.True(t, not_found.Is(m.DeletePost(posts[0].Raw)))
	_, f := m.FindPost("echo")
	assert.FatalOnError(t, f)
}

func TestMockPostsToLookedUpChannelByIdentifier(t *testing.T) {
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
