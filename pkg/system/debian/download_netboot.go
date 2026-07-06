package debian

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/debian/netboot"
	"github.com/funtimecoding/soil/pkg/system/debian/preseed"
	"github.com/funtimecoding/soil/pkg/system/debian/release"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func (c *Client) DownloadNetboot(
	r *release.Release,
	architecture string,
) {
	inputArchive := join.Absolute(c.workDirectory, netboot.Name(r.Codename))
	netboot.Download(r.Codename, architecture, inputArchive)
	inputConfiguration := join.Absolute(
		c.workDirectory,
		preseed.Name(r.Codename),
	)
	preseed.Download(r.Codename, inputConfiguration)
	temporary := join.Absolute(c.workDirectory, "patched_netboot")

	if system.FileExists(temporary) {
		system.Remove(temporary)
	}

	PrepareNetBoot(
		inputArchive,
		join.Absolute(c.workDirectory, "patched_netboot.tar.gz"),
		inputConfiguration,
		temporary,
	)
	system.Remove(temporary)
}
