package filter

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/maps"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (f *Filter) Build() string {
	var parts []string

	for _, k := range maps.StringKeys(f.equal) {
		parts = append(parts, fmt.Sprintf("%s=\"%s\"", k, f.equal[k]))
	}

	for _, k := range maps.StringKeys(f.like) {
		parts = append(parts, fmt.Sprintf("%s=~\"%s\"", k, f.like[k]))
	}

	return fmt.Sprintf("{%s}", join.Comma(parts))
}
