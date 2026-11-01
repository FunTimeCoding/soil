package chip

import "maragu.dev/gomponents"

type Chip struct {
	label      string
	link       string
	kinds      []string
	color      string
	toggle     bool
	active     bool
	current    bool
	attributes []gomponents.Node
	content    []gomponents.Node
}
