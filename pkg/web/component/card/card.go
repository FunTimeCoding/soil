package card

import "maragu.dev/gomponents"

type Card struct {
	name       string
	link       string
	portrait   string
	emblem     string
	treatment  string
	size       string
	task       bool
	preserve   bool
	accent     string
	marker     gomponents.Node
	trailing   []gomponents.Node
	subtitle   string
	meta       []string
	state      []gomponents.Node
	footer     []gomponents.Node
	identifier string
	terms      string
	class      string
}
