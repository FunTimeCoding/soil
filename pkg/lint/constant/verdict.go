package constant

const (
	VerdictLive Verdict = iota
	VerdictDead
	VerdictAbsolute
	VerdictConvention
	VerdictBareSlash
	VerdictUndeclaredHost
	VerdictTallied
)

type Verdict int
