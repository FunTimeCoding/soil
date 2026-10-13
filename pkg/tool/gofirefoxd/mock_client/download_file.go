package mock_client

import "github.com/funtimecoding/soil/pkg/firefox/download"

func (c *Client) DownloadFile(
	l string,
	path string,
) (*download.Download, error) {
	c.downloadIdentifier++
	result := download.New()
	result.Identifier = c.downloadIdentifier
	result.Locator = l
	result.Path = path
	result.State = "in_progress"
	c.downloads = append(c.downloads, result)

	return result, nil
}
