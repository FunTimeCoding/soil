package main

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/generative/n8n"
)

func main() {
	// https://docs.n8n.io/api/api-reference/
	for _, w := range n8n.NewEnvironment().Workflows() {
		console.Format("Workflow: %s\n", w.Name)

		for _, n := range w.Nodes {
			console.Format("  Node: %s\n", n.Name)

			for k, v := range n.Credentials {
				console.Format("    Credential: %s = %s\n", k, v)
			}
		}
	}
}
