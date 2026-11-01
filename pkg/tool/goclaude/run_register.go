package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func RunRegister(
	c face.Client,
	session string,
) string {
	r, e := c.PostRegister(
		context.Background(),
		client.PostRegisterJSONRequestBody{Session: session},
	)

	if e != nil {
		return ""
	}

	return r.Callsign
}
