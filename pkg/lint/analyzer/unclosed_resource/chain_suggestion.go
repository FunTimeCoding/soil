package unclosed_resource

import "go/types"

func chainSuggestion(t types.Type) string {
	if returnsError(t) {
		return "defer errors.PanicClose on it"
	}

	return "defer its Close"
}
