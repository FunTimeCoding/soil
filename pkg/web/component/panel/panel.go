package panel

import "maragu.dev/gomponents"

type Panel struct {
	title      string
	identifier string
	class      string
	header     []gomponents.Node
	content    []gomponents.Node
	footer     []gomponents.Node
}
