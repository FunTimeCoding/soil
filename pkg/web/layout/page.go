package layout

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"maragu.dev/gomponents"
)

type Page struct {
	identity     *identity.Tool
	title        string
	path         string
	theme        string
	favicon      []byte
	style        string
	items        []*navigation_item.Item
	navigation   []gomponents.Node
	script       []string
	footer       []gomponents.Node
	content      []gomponents.Node
	summary      []string
	liveEndpoint string
	liveParams   string
	commands     []*palette.Command
	extended     string
	serverSide   string
	fullWindow   bool
}
