package field

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func (f *Field) suggestionList() string {
	if len(f.suggestions) == 0 {
		return ""
	}

	return join.Empty(f.name, constant.EntrySuggestionsSuffix)
}
