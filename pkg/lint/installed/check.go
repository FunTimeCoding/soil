package installed

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"golang.org/x/mod/semver"
	"path/filepath"
)

func Check(root string) []*concern.Concern {
	sources := modules(filepath.Dir(root))
	names := commands(sources)
	known := make(map[string]bool, len(names))

	for n := range names {
		known[n] = true
	}

	latest := make(map[string]string)
	var result []*concern.Concern

	for _, b := range PathBinaries(known) {
		module := b.Module

		if _, okay := sources[module]; !okay {
			candidates := names[b.Name]

			if len(candidates) != 1 {
				continue
			}

			module = candidates[0]
		}

		directory := sources[module]

		if !semver.IsValid(b.Version) {
			result = append(
				result,
				finding(
					constant.UnresolvedBinaryKey,
					constant.UnresolvedBinaryText,
					directory,
					b,
				),
			)

			continue
		}

		if _, okay := latest[directory]; !okay {
			latest[directory] = latestTag(directory)
		}

		if Stale(b.Version, latest[directory]) {
			result = append(
				result,
				finding(
					constant.StaleBinaryKey,
					constant.StaleBinaryText,
					directory,
					b,
				),
			)
		}
	}

	return result
}
