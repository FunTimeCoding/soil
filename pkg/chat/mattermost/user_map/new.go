package user_map

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/mattermost/mattermost/server/public/model"
)

func New(
	userAlias face.StringAlias,
	administrator []string,
) *Map {
	return &Map{
		byName:        make(map[string]*model.User),
		byIdentifier:  make(map[string]*model.User),
		administrator: administrator,
		userAlias:     userAlias,
	}
}
