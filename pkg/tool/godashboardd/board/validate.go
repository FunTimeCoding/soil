package board

import "fmt"

func (b *Board) validate() {
	seen := map[string]bool{}

	for _, v := range b.Entries() {
		if v.Label == "" {
			panic("entry without label")
		}

		if seen[v.Label] {
			panic(fmt.Sprintf("duplicate entry label: %s", v.Label))
		}

		seen[v.Label] = true
	}
}
