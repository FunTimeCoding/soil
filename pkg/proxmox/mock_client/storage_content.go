package mock_client

import (
	"fmt"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) StorageContent(s *proxmox.Storage) ([]*proxmox.StorageContent, error) {
	key := fmt.Sprintf("%s/%s", s.Node, s.Name)

	return c.storageContent[key], nil
}
