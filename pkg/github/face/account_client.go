package face

import (
	"github.com/funtimecoding/soil/pkg/github/user"
	"github.com/google/go-github/v93/github"
)

type AccountClient interface {
	MustUser() *user.User
	Notifications() ([]*github.Notification, error)
	Organizations(user string) ([]*github.Organization, error)
	User() (*user.User, error)
}
