package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"testing"
)

func writeEach(
	t *testing.T,
	m *mock_client.Client,
) []error {
	t.Helper()
	x := t.Context()
	var result []error
	add := func(e error) {
		if e != nil {
			result = append(result, e)
		}
	}
	_, e := m.PostIndex(
		x,
		client.PostIndexJSONRequestBody{Collection: new("alfa")},
	)
	add(e)
	_, e = m.PostEmbed(x)
	add(e)
	_, e = m.PostRechunk(x)
	add(e)
	add(m.PostTag(x, client.PostTagJSONRequestBody{Path: "bravo.md"}))
	add(m.PostCollection(x, client.PostCollectionJSONRequestBody{Name: "alfa"}))
	_, e = m.DeleteCollection(x, &client.DeleteCollectionParams{Name: "alfa"})
	add(e)
	add(
		m.PostContext(
			x,
			client.PostContextJSONRequestBody{
				Collection: "alfa",
				PathPrefix: "charlie",
			},
		),
	)
	_, e = m.DeleteContext(
		x,
		&client.DeleteContextParams{Collection: "alfa", PathPrefix: "charlie"},
	)
	add(e)

	return result
}
