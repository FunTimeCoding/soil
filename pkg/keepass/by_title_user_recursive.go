package keepass

import (
	"github.com/funtimecoding/soil/pkg/keepass/constant"
	"github.com/tobischo/gokeepasslib/v3"
)

func ByTitleUserRecursive(
	g []gokeepasslib.Group,
	title string,
	user string,
) *gokeepasslib.Entry {
	for _, group := range g {
		for _, entry := range group.Entries {
			if entry.GetTitle() == title {
				if v := entry.Get(constant.UserNameKey); v != nil && v.Value.Content == user {
					return &entry
				}
			}
		}

		if e := ByTitleUserRecursive(group.Groups, title, user); e != nil {
			return e
		}
	}

	return nil
}
