package constant

type Policy int

const (
	Pessimistic Policy = iota
	Optimistic
	Skip
)
