package grant

import "slices"

func (g *Grant) Holds(repository string) bool {
	g.load()

	return slices.Contains(g.repositories, repository)
}
