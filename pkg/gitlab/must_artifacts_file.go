package gitlab

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustArtifactsFile(
	project int64,
	reference string,
	job string,
) string {
	result, e := c.ArtifactsFile(project, reference, job)
	errors.PanicOnError(e)

	return result
}
