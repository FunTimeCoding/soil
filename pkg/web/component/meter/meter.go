package meter

import "maragu.dev/gomponents"

type Meter struct {
	value   float64
	maximum float64
	label   []gomponents.Node
	shown   string
	color   string
	track   string
	width   string
	height  string
	icon    gomponents.Node
	title   string
	class   string
	gauge   bool
}
