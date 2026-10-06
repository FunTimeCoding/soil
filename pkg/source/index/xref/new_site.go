package xref

func NewSite(
	file string,
	line int,
	column int,
) *Site {
	return &Site{File: file, Line: line, Column: column}
}
