package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func RunTurnEnd(
	c face.Client,
	session string,
) {
	if e := c.PostTurnEnd(
		context.Background(),
		client.PostTurnEndJSONRequestBody{Session: session},
	); e != nil {
		return
	}
}
