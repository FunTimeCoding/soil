package action

import "maragu.dev/gomponents"

type Action struct {
	path        string
	label       string
	title       string
	fields      []gomponents.Node
	target      string
	swap        string
	confirm     string
	disabled    bool
	small       bool
	class       string
	buttonClass string
}
