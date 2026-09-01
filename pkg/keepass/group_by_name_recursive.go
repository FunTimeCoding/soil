package keepass

import "github.com/tobischo/gokeepasslib/v3"

func GroupByNameRecursive(
	g []gokeepasslib.Group,
	name string,
) *gokeepasslib.Group {
	for i := range g {
		if g[i].Name == name {
			return &g[i]
		}

		if found := GroupByNameRecursive(g[i].Groups, name); found != nil {
			return found
		}
	}

	return nil
}
