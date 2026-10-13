package example

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/funtimecoding/soil/pkg/console"
)

func Team() {
	m := mattermost.NewEnvironment()
	defer m.Close()

	for _, t := range m.MustTeams(m.MustMe().Id) {
		console.Format("Team: %s %s\n", t.Id, t.Name)
	}
}
