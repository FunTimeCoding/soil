package layout

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"maragu.dev/gomponents"
)

type Page struct {
	identity        *identity.Tool
	title           string
	path            string
	theme           string
	style           string
	brandNode       gomponents.Node
	items           []*navigation_item.Item
	navigation      []gomponents.Node
	script          []string
	footer          []gomponents.Node
	content         []gomponents.Node
	summary         []string
	liveEndpoint    string
	liveParams      string
	paletteEndpoint string
}
