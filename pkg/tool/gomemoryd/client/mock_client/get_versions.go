package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/version_entry"
)

func (c *Client) GetVersions(
	context.Context,
	string,
	int,
) ([]version_entry.Entry, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return []version_entry.Entry{}, nil
}
