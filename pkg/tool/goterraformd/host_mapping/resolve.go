package host_mapping

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"strings"
)

func (x *Index) resolve(
	prefix string,
	reference string,
) []string {
	s := strings.Split(
		configurationAddress(reference),
		constant.ModuleSeparator,
	)

	switch {
	case s[0] == constant.VariableScope && len(s) > 1:
		return x.resolveVariable(prefix, s[1])
	case s[0] == constant.ModuleScope && len(s) > 2:
		return x.resolveOutput(prefix, s[1], s[2])
	case s[0] == constant.SourceScope && len(s) > 2:
		return x.existing(join.Empty(prefix, join.Dot(s[:3])))
	case len(s) > 1:
		return x.existing(join.Empty(prefix, join.Dot(s[:2])))
	default:
		return nil
	}
}
