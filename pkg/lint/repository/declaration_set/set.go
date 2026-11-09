package declaration_set

type Set struct {
	Names   map[string]bool
	Types   map[string]bool
	members map[string]map[string]bool
	embeds  map[string][]string
	foreign map[string]bool
}
