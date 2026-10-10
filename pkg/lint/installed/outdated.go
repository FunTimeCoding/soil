package installed

import (
	"golang.org/x/mod/semver"
	"path/filepath"
)

func Outdated(root string) []*Binary {
	sources := modules(filepath.Dir(root))
	names := commands(sources)
	known := make(map[string]bool, len(names))

	for n := range names {
		known[n] = true
	}

	latest := make(map[string]string)
	var result []*Binary

	for _, b := range PathBinaries(known) {
		module := b.Module

		if _, okay := sources[module]; !okay {
			candidates := names[b.Name]

			if len(candidates) != 1 {
				continue
			}

			module = candidates[0]
		}

		b.Directory = sources[module]

		if !semver.IsValid(b.Version) {
			result = append(result, b)

			continue
		}

		if _, okay := latest[b.Directory]; !okay {
			latest[b.Directory] = latestTag(b.Directory)
		}

		if Stale(b.Version, latest[b.Directory]) {
			result = append(result, b)
		}
	}

	return result
}
