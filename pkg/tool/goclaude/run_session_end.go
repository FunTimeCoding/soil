package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func RunSessionEnd(
	c face.Client,
	session string,
	reason string,
) {
	body := client.PostSessionEndJSONRequestBody{Session: session}

	if reason != "" {
		body.Reason = &reason
	}

	if e := c.PostSessionEnd(context.Background(), body); e != nil {
		return
	}
}
