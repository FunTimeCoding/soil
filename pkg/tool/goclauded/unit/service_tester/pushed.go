package service_tester

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"

func (t *Tester) Pushed(collection string) []client.PostDocumentJSONRequestBody {
	var result []client.PostDocumentJSONRequestBody

	for _, p := range t.Query.Pushed {
		if p.Collection == collection {
			result = append(result, p)
		}
	}

	return result
}
