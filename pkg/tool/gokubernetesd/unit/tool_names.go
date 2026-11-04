package unit

import "github.com/funtimecoding/soil/pkg/generative/model_context_client"

func toolNames(c *model_context_client.Client) []string {
	var result []string

	for _, tool := range c.ListTools() {
		result = append(result, tool.Name)
	}

	return result
}
