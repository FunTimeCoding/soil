package option

type Analyze struct {
	Root     string
	Summary  bool
	Comment  bool
	Full     bool
	Index    string
	Patterns []string
}
