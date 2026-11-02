package toolchain

import "github.com/funtimecoding/soil/pkg/tool/gosourced/types/removal"

func logValues() []*removal.Parameter {
	return []*removal.Parameter{
		removal.NewParameter("other.test/lib", "Log", "", []string{"values"}),
	}
}
