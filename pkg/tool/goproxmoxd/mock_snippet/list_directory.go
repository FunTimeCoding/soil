package mock_snippet

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/mock_snippet/mock_entry"
	"io/fs"
)

func (c *Client) ListDirectory(_ string) []fs.FileInfo {
	var result []fs.FileInfo

	for name, content := range c.files {
		result = append(
			result,
			mock_entry.New(name, int64(len(content))),
		)
	}

	return result
}
