package gofix

type formatChange struct {
	Kind    string
	Message string
	Offset  int
	Line    int
}
