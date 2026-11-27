package tool_case

import "github.com/funtimecoding/soil/pkg/browser/firefox/mock_client"

type Tool struct {
	Name      string
	Arguments map[string]any
	Method    func(*mock_client.Client) any
}
