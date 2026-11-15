package peer_tester

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"

func (s *Peers) Pushed(collection string) []client.PostDocumentJSONRequestBody {
	var result []client.PostDocumentJSONRequestBody

	for _, p := range s.Query.Pushed {
		if p.Collection == collection {
			result = append(result, p)
		}
	}

	return result
}
