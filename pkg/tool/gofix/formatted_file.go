package gofix

type formattedFile struct {
	Source    []byte
	Changes   []*formatChange
	Converged bool
}
