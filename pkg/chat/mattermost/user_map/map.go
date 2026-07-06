package user_map

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/mattermost/mattermost/server/public/model"
)

type Map struct {
	byName        map[string]*model.User
	byIdentifier  map[string]*model.User
	administrator []string
	userAlias     face.StringAlias
}
