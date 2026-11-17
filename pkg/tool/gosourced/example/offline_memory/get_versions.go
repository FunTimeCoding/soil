package offline_memory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/version_entry"
)

func (c *Client) GetVersions(
	context.Context,
	string,
	int,
) ([]version_entry.Entry, error) {
	return nil, nil
}
