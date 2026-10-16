package host_mapping

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
)

func (x *Index) resolveOutput(
	prefix string,
	module string,
	name string,
) []string {
	child := join.Empty(
		prefix,
		constant.ModuleScope,
		constant.ModuleSeparator,
		module,
		constant.ModuleSeparator,
	)
	o, okay := x.outputs[child][name]

	if !okay || o == nil {
		return nil
	}

	var result []string

	for _, r := range references(o.Expression) {
		result = append(result, x.resolve(child, r)...)
	}

	return result
}
