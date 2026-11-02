package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) DownloadLocator(
	_ *proxmox.Storage,
	_ string,
	_ string,
	_ string,
) (*proxmox.Task, error) {
	return &proxmox.Task{UPID: "mock:download"}, nil
}
