package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
	"testing"
)

func TestPost(t *testing.T) {
	assert.NotNil(t, post.New(&model.Post{}))
}
