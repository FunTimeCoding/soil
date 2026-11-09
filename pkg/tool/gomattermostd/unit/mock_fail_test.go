package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/client"
	"testing"
	"time"
)

func TestMockReactionsFailOnTheirOwn(t *testing.T) {
	m := mock_client.New()
	t.Cleanup(m.Close)
	m.Reactions["alfa"] = []string{"bravo"}
	m.Posts["charlie"] = []*client.Post{{Identifier: "alfa"}}
	m.Fail(m.GetPostReactions, unavailable.Format("gomattermostd: down"))
	_, e := m.GetPostReactions(context.Background(), "alfa")
	assert.True(t, unavailable.Is(e))
	posts, f := m.GetChannelPosts(
		context.Background(),
		"charlie",
		&client.GetChannelPostsParams{Since: time.Time{}},
	)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, posts)
}
