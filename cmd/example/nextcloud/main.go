package main

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/nextcloud"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func main() {
	n := nextcloud.NewEnvironment()

	if false {
		n.MustStatus()
		console.Line("success")
	}

	if false {
		for _, f := range n.MustReadDirectory("/") {
			console.Format(
				"Name: %s, IsDir: %t, Size: %d\n",
				f.Name(),
				f.IsDir(),
				f.Size(),
			)
		}
	}

	if false {
		n.MustDownloadFile(
			"example.png",
			join.Absolute(system.Home(), constant.DownloadsPath, "example.png"),
		)
	}

	if false {
		n.MustUploadFile(
			"example2.png",
			join.Absolute(system.Home(), constant.DownloadsPath),
			"example.png",
		)
	}
}
