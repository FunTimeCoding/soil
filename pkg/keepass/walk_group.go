package keepass

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/tobischo/gokeepasslib/v3"
)

func walkGroup(
	prefix string,
	group *gokeepasslib.Group,
	visit func(
		path string,
		group *gokeepasslib.Group,
		entry *gokeepasslib.Entry,
	),
) {
	path := group.Name

	if prefix != "" {
		path = join.Empty(prefix, "/", group.Name)
	}

	for i := range group.Entries {
		visit(path, group, &group.Entries[i])
	}

	for i := range group.Groups {
		walkGroup(path, &group.Groups[i], visit)
	}
}
