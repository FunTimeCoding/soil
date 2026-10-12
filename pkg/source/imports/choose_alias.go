package imports

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

	for _, c := range []string{s, Fallback(s, declaredName)} {
		if !taken(c) {
			return c
		}
	}

	return ""
}
