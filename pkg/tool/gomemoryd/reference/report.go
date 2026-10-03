package reference

type Report struct {
	Identifier int64
	Name       string
	Findings   []*Finding
}
