package option

import "io"

type Analyze struct {
	Root     string
	Summary  bool
	Comment  bool
	Full     bool
	Index    string
	Patterns []string
	Notice   io.Writer
}
