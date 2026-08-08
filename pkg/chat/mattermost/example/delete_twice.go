package example

import (
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/mattermost/mattermost/server/public/model"
)

func DeleteTwice() {
	a := argument.NewSimple("mattermost-delete-twice")
	a.ParseSimple()
	channel := a.RequiredPositional(0, "CHANNEL")
	m := mattermost.NewEnvironment(mattermost.WithVerbose(false))
	c := m.MustTeamChannel(channel)
	p := m.MustPostSimple(c, "delete-twice probe")
	m.MustDeletePost(p)
	e := m.DeletePost(p)
	fmt.Printf("Second delete error: %v\n", e)
	var f *model.AppError

	if errors.As(e, &f) {
		fmt.Printf(
			"AppError identifier %s status %d\n",
			f.Id,
			f.StatusCode,
		)
	}
}
