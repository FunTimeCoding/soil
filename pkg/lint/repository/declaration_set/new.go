package declaration_set

func New() *Set {
	return &Set{
		Names:   map[string]bool{},
		Types:   map[string]bool{},
		members: map[string]map[string]bool{},
		embeds:  map[string][]string{},
		foreign: map[string]bool{},
	}
}
