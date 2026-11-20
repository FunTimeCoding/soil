package overview_result

type Result struct {
	Rows      []*Row
	Open      int
	Critical  int
	DebianFix int
	OwnWork   int
}
