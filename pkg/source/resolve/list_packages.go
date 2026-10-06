package resolve

import (
	"fmt"
	"golang.org/x/tools/go/packages"
)

func ListPackages(
	directory string,
	patterns ...string,
) (map[string]bool, error) {
	c := &packages.Config{
		Mode:       packages.NeedName,
		Dir:        directory,
		Tests:      true,
		BuildFlags: BuildFlags(directory),
	}
	listed, e := packages.Load(c, patterns...)

	if e != nil {
		return nil, fmt.Errorf("list packages: %w", e)
	}

	result := make(map[string]bool)

	for _, p := range listed {
		result[p.PkgPath] = true
	}

	return result, nil
}
