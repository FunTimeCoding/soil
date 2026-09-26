package aptly

import "github.com/funtimecoding/soil/pkg/semver"

func (c *Client) LatestVersion(
	repository string,
	name string,
) (string, error) {
	packages, e := c.Packages(repository)

	if e != nil {
		return "", e
	}

	versions := PackageVersions(packages, name)

	if len(versions) == 0 {
		return "", nil
	}

	semver.SortDescending(versions)

	return versions[0], nil
}
