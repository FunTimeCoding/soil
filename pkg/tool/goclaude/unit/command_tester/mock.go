package command_tester

import "github.com/funtimecoding/soil/pkg/tool/goclauded/client/mock_client"

func (o *Tester) Mock() *mock_client.Client {
	o.t.Helper()
	result := mock_client.New(o.t)
	result.Resolved = map[string]string{"alfa": "alfa-identifier"}

	return result
}
