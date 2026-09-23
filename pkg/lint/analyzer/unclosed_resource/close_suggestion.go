package unclosed_resource

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/types"
)

func closeSuggestion(o *types.Var) string {
	target := o.Name()

	if isResponse(o.Type()) {
		target = join.Empty(target, ".Body")
	}

	if returnsError(o.Type()) {
		return fmt.Sprintf("defer errors.PanicClose(%s)", target)
	}

	return fmt.Sprintf("defer %s.Close()", target)
}
