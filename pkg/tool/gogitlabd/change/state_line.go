package change

import "github.com/funtimecoding/soil/pkg/strings/join"

func StateLine(
	state string,
	actor string,
) string {
	if actor == "" {
		return state
	}

	return join.Empty(state, " by ", actor)
}
