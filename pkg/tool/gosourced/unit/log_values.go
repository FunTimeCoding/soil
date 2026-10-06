package unit

import "github.com/funtimecoding/soil/pkg/tool/gosourced/service/removal"

func logValues() []*removal.Parameter {
	return []*removal.Parameter{
		removal.NewParameter("other.test/lib", "Log", "", []string{"values"}),
	}
}
