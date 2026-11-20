package command_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (o *Tester) Mock() *mock_client.Client {
	o.t.Helper()
	result := mock_client.New(o.t)
	_, e := result.PostDocument(
		o.t.Context(),
		client.PostDocumentJSONRequestBody{
			Collection: "alfa",
			Path:       "kilo.md",
			Body:       "lima",
		},
	)
	assert.FatalOnError(o.t, e)
	result.ContextsAdded = append(
		result.ContextsAdded,
		client.PostContextJSONRequestBody{
			Collection: "alfa",
			PathPrefix: "kilo",
		},
	)

	return result
}
