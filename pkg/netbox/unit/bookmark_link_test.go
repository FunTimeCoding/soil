package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/bookmark"
	"github.com/funtimecoding/soil/pkg/netbox/unit/netbox_tester"
	"testing"
)

func TestBookmarkLinkDropsInterfaceSegment(t *testing.T) {
	raw := netbox_tester.Bookmark()
	raw.Object = map[string]any{
		"url": "https://host.example/api/dcim/devices/2132/",
	}
	result := bookmark.New(raw)
	assert.String(t, "https://host.example/dcim/devices/2132/", result.Link)
	assert.String(t, "alfa", result.Display)
	assert.String(t, "dcim.device", result.ObjectType)
}

func TestBookmarkLinkEmptyWithoutObjectLink(t *testing.T) {
	assert.String(t, "", bookmark.New(netbox_tester.Bookmark()).Link)
}

func TestBookmarkLinkEmptyWhenObjectCarriesNoLink(t *testing.T) {
	raw := netbox_tester.Bookmark()
	raw.Object = map[string]any{"display": "alfa"}
	assert.String(t, "", bookmark.New(raw).Link)
}
