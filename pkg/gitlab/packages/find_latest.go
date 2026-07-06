package packages

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"log"
)

func FindLatest(v []*gitlab.Package) *gitlab.Package {
	if len(v) == 0 {
		return nil
	}

	latest := Latest(v)

	if len(latest) == 1 {
		return latest[v[0].Name]
	}

	var names []string

	for _, p := range latest {
		names = append(names, p.Name)
	}

	log.Panicf(
		"multiple latest versions found for packages: %s",
		join.Comma(names),
	)

	return nil
}
