package client

import (
	"context"
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"io"
)

func (c *RestClient) Statistics() *client.Statistics {
	r, e := c.http.GetStatistics(context.Background())

	if e != nil {
		return nil
	}

	defer errors.PanicClose(r.Body)
	body, e := io.ReadAll(r.Body)

	if e != nil {
		return nil
	}

	var result client.Statistics

	if json.Unmarshal(body, &result) != nil {
		return nil
	}

	return &result
}
