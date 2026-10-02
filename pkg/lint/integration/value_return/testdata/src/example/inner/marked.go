package inner

//goanalyze:by-value measured - heap allocated on every step
type Marked struct {
	Value int
}

//goanalyze:by-value
type BareDirective struct {
	Value int
}

type Unmarked struct {
	Value int
}
