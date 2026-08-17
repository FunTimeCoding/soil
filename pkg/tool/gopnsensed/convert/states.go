package convert

import (
	"github.com/funtimecoding/soil/pkg/opnsense/state"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func States(v []*state.State) []server.State {
	result := []server.State{}

	for _, e := range v {
		result = append(result, *State(e))
	}

	return result
}
