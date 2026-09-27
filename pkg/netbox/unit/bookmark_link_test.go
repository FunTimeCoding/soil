package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/bookmark"
	"testing"
)

func TestBookmarkLinkDropsInterfaceSegment(t *testing.T) {
	raw := newBookmark()
	raw.Object = map[string]any{
		"url": "https://host.example/api/dcim/devices/2132/",
	}
	result := bookmark.New(raw)
	assert.String(t, "https://host.example/dcim/devices/2132/", result.Link)
	assert.String(t, "alfa", result.Display)
	assert.String(t, "dcim.device", result.ObjectType)
}

func TestBookmarkLinkEmptyWithoutObjectLink(t *testing.T) {
	assert.String(t, "", bookmark.New(newBookmark()).Link)
}

func TestBookmarkLinkEmptyWhenObjectCarriesNoLink(t *testing.T) {
	raw := newBookmark()
	raw.Object = map[string]any{"display": "alfa"}
	assert.String(t, "", bookmark.New(raw).Link)
}
