package rotation

import "github.com/funtimecoding/soil/pkg/face"

func ContainsUser(
	r *Rotation,
	user string,
	f face.StringAlias,
) bool {
	for _, e := range r.Participants {
		if f == nil {
			if e.Username == user {
				return true
			}
		} else {
			if f(e.Username) == user {
				return true
			}
		}
	}

	return false
}
