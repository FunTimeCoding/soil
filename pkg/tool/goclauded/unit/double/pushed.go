package double

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"

func (s *Set) Pushed(collection string) []client.PostDocumentJSONRequestBody {
	var result []client.PostDocumentJSONRequestBody

	for _, p := range s.Query.Pushed {
		if p.Collection == collection {
			result = append(result, p)
		}
	}

	return result
}
