package system

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/installed"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
	"github.com/funtimecoding/soil/pkg/vulnerability/debian/tracker"
)

func matchAll(
	t *tracker.Tracker,
	release string,
	packages []installed.Installed,
	supported bool,
) []*system_result.Finding {
	var result []*system_result.Finding
	seen := map[[2]string]int{}

	for _, p := range packages {
		if kernel(p.Source) {
			continue
		}

		for _, f := range match(t, release, &p, supported) {
			k := [2]string{f.Vulnerability, f.Source}
			i, okay := seen[k]

			if !okay {
				seen[k] = len(result)
				result = append(result, f)

				continue
			}

			if f.Upgradable && !result[i].Upgradable {
				result[i] = f
			}
		}
	}

	return result
}
