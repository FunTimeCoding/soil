package post

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/mattermost/mattermost/server/public/model"
	"testing"
)

func TestFromList(t *testing.T) {
	bravo := &model.Post{Id: "bravo", CreateAt: 3000}
	alfa := &model.Post{Id: "alfa", CreateAt: 2000}
	parent := &model.Post{Id: "parent", CreateAt: 1000}
	l := &model.PostList{
		Order: []string{"bravo", "alfa", "missing"},
		Posts: map[string]*model.Post{
			"alfa":   alfa,
			"bravo":  bravo,
			"parent": parent,
		},
	}
	assert.Any(t, []*model.Post{bravo, alfa}, FromList(l, false))
	assert.Any(t, []*model.Post{alfa, bravo}, FromList(l, true))
}
