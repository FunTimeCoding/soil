package aptly

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/semver"
	"strings"
)

func (c *Client) LatestVersion(
	repository string,
	name string,
) (string, error) {
	packages, e := c.Packages(repository)

	if e != nil {
		return "", e
	}

	prefix := fmt.Sprintf("%s_", name)
	var versions []string

	for _, p := range packages {
		if !strings.HasPrefix(p, prefix) {
			continue
		}

		parts := strings.SplitN(p, "_", 3)

		if len(parts) < 2 {
			continue
		}

		versions = append(versions, parts[1])
	}

	if len(versions) == 0 {
		return "", nil
	}

	semver.SortDescending(versions)

	return versions[0], nil
}
