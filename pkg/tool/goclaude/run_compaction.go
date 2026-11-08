package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func RunCompaction(
	c face.Client,
	session string,
	phase client.CompactionPhase,
	trigger string,
) {
	body := client.PostCompactionJSONRequestBody{
		Session: session,
		Phase:   phase,
	}

	if trigger != "" {
		body.Trigger = &trigger
	}

	if e := c.PostCompaction(context.Background(), body); e != nil {
		return
	}
}
