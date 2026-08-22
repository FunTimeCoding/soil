package gitlab

import "bytes"

func (c *Client) ArtifactFile(
	project int64,
	job int64,
	path string,
) (*bytes.Reader, error) {
	result, _, e := c.client.Jobs.DownloadSingleArtifactsFile(
		project,
		job,
		path,
	)

	return result, wrapError(e)
}
