package imports

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"slices"
	"strings"
)

func Keeper(
	packagePath string,
	declaredName string,
	importPaths []string,
	plain string,
	isStandard func(string) bool,
) (string, string) {
	if IsVersion(declaredName) {
		return "", constant.KeeperVersion
	}

	var standard []string

	for _, p := range importPaths {
		if isStandard(p) {
			standard = append(standard, p)
		}
	}

	if len(standard) == 1 {
		return standard[0], constant.KeeperStandard
	}

	if own := Own(packagePath, importPaths); own != "" {
		return own, constant.KeeperOwn
	}

	deepest := -1
	var deep []string

	for _, p := range importPaths {
		n := strings.Count(p, "/")

		if n > deepest {
			deepest = n
			deep = []string{p}

			continue
		}

		if n == deepest {
			deep = append(deep, p)
		}
	}

	if len(deep) == 1 {
		return deep[0], constant.KeeperDeeper
	}

	if slices.Contains(deep, plain) {
		return plain, constant.KeeperPlain
	}

	return slices.Min(deep), constant.KeeperFirst
}
