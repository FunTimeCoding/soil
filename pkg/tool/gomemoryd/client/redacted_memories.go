package client

import (
	"context"
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"io"
)

func (c *RestClient) RedactedMemories() map[int64]bool {
	r, e := c.http.GetRedactedMemories(context.Background())

	if e != nil {
		return nil
	}

	defer errors.PanicClose(r.Body)
	body, e := io.ReadAll(r.Body)

	if e != nil {
		return nil
	}

	var identifiers []int64

	if json.Unmarshal(body, &identifiers) != nil {
		return nil
	}

	result := map[int64]bool{}

	for _, identifier := range identifiers {
		result[identifier] = true
	}

	return result
}
