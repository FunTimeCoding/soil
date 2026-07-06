package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/override"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/rotation"
)

func UserInvolved(
	user string,
	o []*override.Override,
	r []*rotation.Rotation,
) bool {
	for _, e := range o {
		if override.ContainsUser(e, user, nil) {
			return true
		}
	}

	for _, e := range r {
		if rotation.ContainsUser(e, user, nil) {
			return true
		}
	}

	return false
}
