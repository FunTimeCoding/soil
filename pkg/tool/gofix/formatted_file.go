package gofix

type FormattedFile struct {
	Source    []byte
	Changes   []*FormatChange
	Converged bool
}
