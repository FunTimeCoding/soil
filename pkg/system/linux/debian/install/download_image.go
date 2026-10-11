package install

import (
	"github.com/funtimecoding/soil/pkg/system/linux/debian/install/image"
	"github.com/funtimecoding/soil/pkg/system/linux/debian/install/image/checksum"
	"github.com/funtimecoding/soil/pkg/system/linux/debian/install/release"
)

func (c *Client) DownloadImage(
	r *release.Release,
	architecture string,
) {
	image.Download(r.Version(), architecture, c.workDirectory)
	checksum.Download(r.Version(), architecture, c.workDirectory)
}
