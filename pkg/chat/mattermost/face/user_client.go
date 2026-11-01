package face

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/user_map"
	"github.com/mattermost/mattermost/server/public/model"
)

type UserClient interface {
	Me() (*model.User, error)
	MustMe() *model.User
	MustUser(identifier string) *model.User
	MustUsers(
		page int,
		perPage int,
	) []*model.User
	SearchUsers(
		t *model.Team,
		query string,
		limit int,
	) ([]*model.User, error)
	User(identifier string) (*model.User, error)
	UserByName(name string) (*model.User, error)
	UserMap() *user_map.Map
	Users(
		page int,
		perPage int,
	) ([]*model.User, error)
}
