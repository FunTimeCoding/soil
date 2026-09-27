package unit

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/netbox-community/go-netbox/v4"
)

func newBookmark() *netbox.Bookmark {
	return netbox.NewBookmark(
		1,
		"https://host.example/api/extras/bookmarks/1/",
		"alfa",
		constant.DeviceAddress,
		2132,
		*netbox.NewBriefUser(
			1,
			"https://host.example/api/users/users/1/",
			"bravo",
			"bravo",
		),
	)
}
