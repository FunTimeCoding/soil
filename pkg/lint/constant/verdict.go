package constant

const (
	VerdictLive Verdict = iota
	VerdictDead
	VerdictAbsolute
	VerdictConvention
	VerdictBareSlash
	VerdictUndeclaredHost
	VerdictTallied
	VerdictDeadHeading
	VerdictFragmentTarget
)

type Verdict int
