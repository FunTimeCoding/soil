package constant

import "regexp"

const (
	TestPlaceholder = "Xxx"
	CallSuffix      = "()"
)

var (
	TestSpan        = regexp.MustCompile(`^Test[A-Z][A-Za-z0-9_]*$`)
	TestDeclaration = regexp.MustCompile(`(?m)^func (Test[A-Za-z0-9_]*)\(`)
	BareSymbolSpan  = regexp.MustCompile(
		`^([a-z][a-z0-9_]*\.[A-Z][A-Za-z0-9_]*(\.[A-Za-z_][A-Za-z0-9_]*)?|[A-Z][A-Za-z0-9_]*\.[A-Za-z_][A-Za-z0-9_]*)$`,
	)
)
