package row

import "maragu.dev/gomponents"

type Row struct {
	Label   string
	Value   gomponents.Node
	Tooltip string
	Section bool
}
