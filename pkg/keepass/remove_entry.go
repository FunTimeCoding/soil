package keepass

import (
	"fmt"
	"github.com/tobischo/gokeepasslib/v3"
)

func RemoveEntry(
	group *gokeepasslib.Group,
	identifier string,
) bool {
	for i := range group.Entries {
		if fmt.Sprintf("%x", group.Entries[i].UUID) == identifier {
			group.Entries = append(
				group.Entries[:i],
				group.Entries[i+1:]...,
			)

			return true
		}
	}

	return false
}
