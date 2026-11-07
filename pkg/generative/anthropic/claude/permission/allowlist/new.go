package allowlist

func New(
	words []string,
	prefix bool,
) *Rule {
	return &Rule{Words: words, Prefix: prefix}
}
