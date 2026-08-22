package gitlab

import "bytes"

func (c *Client) Artifacts(
	project int64,
	job int64,
) (*bytes.Reader, error) {
	result, _, e := c.client.Jobs.GetJobArtifacts(project, job)

	return result, wrapError(e)
}
