package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/version_entry"

func (c *Client) VersionsSince(
	_ string,
	_ int,
) []version_entry.Entry {
	return nil
}
