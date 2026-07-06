package post

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/mattermost/mattermost/server/public/model"
	"testing"
)

func TestPost(t *testing.T) {
	assert.NotNil(t, New(&model.Post{}))
}
