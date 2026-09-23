package unclosed_resource

import "go/types"

func isResponse(t types.Type) bool {
	return t.String() == "*net/http.Response"
}
