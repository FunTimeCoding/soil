package mock_client

import (
	"github.com/funtimecoding/soil/pkg/ssh/mock_client/mock_entry"
	"io/fs"
	"path/filepath"
)

func (c *Client) ListDirectory(path string) ([]fs.FileInfo, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	var result []fs.FileInfo

	for name := range c.directories {
		if filepath.Dir(name) == path {
			result = append(
				result,
				mock_entry.NewDirectory(filepath.Base(name)),
			)
		}
	}

	for name, content := range c.files {
		if filepath.Dir(name) != path {
			continue
		}

		result = append(
			result,
			mock_entry.New(filepath.Base(name), int64(len(content))),
		)
	}

	return result, nil
}
