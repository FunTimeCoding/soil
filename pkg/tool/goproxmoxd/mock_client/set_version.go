package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) SetVersion(
	release string,
	repository string,
	version string,
) {
	c.version = &proxmox.Version{
		Release: release,
		RepoID:  repository,
		Version: version,
	}
}
