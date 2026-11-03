package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
	"strings"
)

func (c *Client) EditView(
	_ context.Context,
	identifier int,
	b client.EditViewJSONRequestBody,
) (*client.View, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	if b.OldString == "" {
		return nil, validation.New("old_string is required")
	}

	i, e := c.index(identifier)

	if e != nil {
		return nil, e
	}

	var text string

	if c.ViewList[i].Text != nil {
		text = *c.ViewList[i].Text
	}

	all := b.ReplaceAll != nil && *b.ReplaceAll

	if !all && strings.Count(text, b.OldString) > 1 {
		return nil, conflict.Format(
			"edit view: old_string matches multiple times; use replace_all",
		)
	}

	if all {
		text = strings.ReplaceAll(text, b.OldString, b.NewString)
	} else {
		text = strings.Replace(text, b.OldString, b.NewString, 1)
	}

	c.ViewList[i].Text = nil

	if text != "" {
		c.ViewList[i].Text = new(text)
	}

	c.ViewList[i].IsDirty = true
	result := c.ViewList[i]

	return &result, nil
}
