package constant

const (
	VerdictLive Verdict = iota
	VerdictDead
	VerdictAbsolute
	VerdictConvention
	VerdictBareSlash
	VerdictTallied
)

type Verdict int
