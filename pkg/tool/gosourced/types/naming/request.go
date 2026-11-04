package naming

type Request struct {
	PackagePath string
	Functions   []*Function
	Type        string
	Name        string
	Imports     []string
}
