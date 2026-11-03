package imports

import "github.com/funtimecoding/soil/pkg/source/constant"

func ChooseAlias(
	importPath string,
	modulePath string,
	declaredName string,
	taken func(string) bool,
) string {
	s := Subsystem(importPath, modulePath)

	if s == "" {
		return ""
	}

	candidates := []string{s, Fallback(s, declaredName)}

	if IsLayoutRoot(importPath) && declaredName != constant.ConstantPackage {
		candidates = candidates[1:]
	}

	for _, c := range candidates {
		if !taken(c) {
			return c
		}
	}

	return ""
}
