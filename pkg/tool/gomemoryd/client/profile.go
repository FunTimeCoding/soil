package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"io"
)

func (c *RestClient) Profile(topic string) string {
	params := &client.GetProfileParams{}

	if topic != "" {
		params.Topic = &topic
	}

	r, e := c.http.GetProfile(context.Background(), params)

	if e != nil {
		return ""
	}

	defer errors.PanicClose(r.Body)
	body, e := io.ReadAll(r.Body)

	if e != nil {
		return ""
	}

	return string(body)
}
