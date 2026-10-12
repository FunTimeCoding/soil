package mock_client

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/constant"
	"slices"
	"strings"
)

func (c *Client) List(path string) ([]string, error) {
	prefix := ""

	if path != "" {
		prefix = join.Empty(path, constant.Folder)
	}

	result := []string{}

	for stored := range c.secrets {
		rest, found := strings.CutPrefix(stored, prefix)

		if !found {
			continue
		}

		if head, _, folder := strings.Cut(rest, constant.Folder); folder {
			rest = join.Empty(head, constant.Folder)
		}

		if !slices.Contains(result, rest) {
			result = append(result, rest)
		}
	}

	slices.Sort(result)

	return result, nil
}
