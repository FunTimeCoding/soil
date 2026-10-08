package gofix

type FormatChange struct {
	Kind    string
	Message string
	Offset  int
	Line    int
}
