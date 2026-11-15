package module_tester

import "github.com/funtimecoding/soil/pkg/tool/gosourced/types/removal"

func LogValues() []*removal.Parameter {
	return []*removal.Parameter{
		removal.NewParameter("other.test/lib", "Log", "", []string{"values"}),
	}
}
