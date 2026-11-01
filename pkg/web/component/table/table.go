package table

import "maragu.dev/gomponents"

type Table struct {
	columns        []string
	numeric        map[int]bool
	rows           []gomponents.Node
	empty          string
	identifier     string
	bodyIdentifier string
	class          string
	headerless     bool
}
