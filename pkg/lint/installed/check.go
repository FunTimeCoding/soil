package installed

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"golang.org/x/mod/semver"
)

func Check(root string) []*concern.Concern {
	var result []*concern.Concern

	for _, b := range Outdated(root) {
		if !semver.IsValid(b.Version) {
			result = append(
				result,
				finding(
					constant.UnresolvedBinaryKey,
					constant.UnresolvedBinaryText,
					b,
				),
			)

			continue
		}

		result = append(
			result,
			finding(constant.StaleBinaryKey, constant.StaleBinaryText, b),
		)
	}

	return result
}
