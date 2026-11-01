package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func RunCheck(
	c face.Client,
	session string,
) string {
	r, e := c.GetCheck(
		context.Background(),
		&client.GetCheckParams{Session: session},
	)

	if e != nil {
		return ""
	}

	if !r.Changed {
		return ""
	}

	return r.Context
}
