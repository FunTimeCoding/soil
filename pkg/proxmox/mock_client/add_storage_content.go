package mock_client

import (
	"fmt"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) AddStorageContent(
	node string,
	storage string,
	volid string,
	format string,
	size uint64,
) {
	key := fmt.Sprintf("%s/%s", node, storage)
	c.storageContent[key] = append(
		c.storageContent[key],
		&proxmox.StorageContent{Volid: volid, Format: format, Size: size},
	)
}
