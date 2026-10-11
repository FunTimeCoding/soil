package aptly

import "github.com/funtimecoding/soil/pkg/semver"

func (c *Client) Versions(
	repository string,
	name string,
) ([]string, error) {
	packages, e := c.Packages(repository)

	if e != nil {
		return nil, e
	}

	result := PackageVersions(packages, name)
	semver.SortDescending(result)

	return result, nil
}
