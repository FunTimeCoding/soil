package debian

import (
	"github.com/funtimecoding/soil/pkg/system/debian/image"
	"github.com/funtimecoding/soil/pkg/system/debian/image/checksum"
	"github.com/funtimecoding/soil/pkg/system/debian/release"
)

func (c *Client) DownloadImage(
	r *release.Release,
	architecture string,
) {
	image.Download(r.Version(), architecture, c.workDirectory)
	checksum.Download(r.Version(), architecture, c.workDirectory)
}
