package installed

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	stampConstant "github.com/funtimecoding/soil/pkg/stamp/constant"
	"path/filepath"
)

func Check(root string) []*concern.Concern {
	sources := modules(filepath.Dir(root))
	names := commands(sources)
	known := make(map[string]bool, len(names))

	for n := range names {
		known[n] = true
	}

	directories := make(map[string]map[string]string)
	dependencies := make(map[string]map[string][]string)
	current := make(map[string]map[string]string)
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

		if b.Dirty {
			result = append(
				result,
				finding(
					constant.DirtyBinaryKey,
					constant.DirtyBinaryText,
					directory,
					b,
				),
			)

			continue
		}

		commit := b.Hash

		if commit == "" || commit == stampConstant.DefaultGitHash {
			commit = tagCommit(directory, b.Version)
		}

		if commit == "" {
			continue
		}

		if _, okay := directories[directory]; !okay {
			directories[directory], dependencies[directory] = packages(
				directory,
			)
			current[directory] = versions(directory)
		}

		scope := relevant(
			b.MainPackage(module),
			directories[directory],
			dependencies[directory],
		)

		if len(scope) == 0 {
			continue
		}

		changed, okay := changedFiles(directory, commit)

		if okay && Touches(changed, scope) ||
			Moved(b.Modules, current[directory]) {
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
