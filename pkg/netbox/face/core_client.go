package face

import (
	"github.com/funtimecoding/soil/pkg/netbox/source"
	"github.com/funtimecoding/soil/pkg/netbox/system_number"
	"github.com/funtimecoding/soil/pkg/netbox/user"
	"github.com/funtimecoding/soil/pkg/netbox/user_group"
)

type CoreClient interface {
	MustSources() []*source.Source
	MustSystemNumbers() []*system_number.Number
	MustUserGroups() []*user_group.Group
	MustUsers() []*user.User
	Sources() ([]*source.Source, error)
	SystemNumbers() ([]*system_number.Number, error)
	UserGroups() ([]*user_group.Group, error)
	Users() ([]*user.User, error)
}
