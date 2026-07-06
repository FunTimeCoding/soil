package main

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/n8n"
)

func main() {
	// https://docs.n8n.io/api/api-reference/
	for _, w := range n8n.NewEnvironment().Workflows() {
		fmt.Printf("Workflow: %s\n", w.Name)

		for _, n := range w.Nodes {
			fmt.Printf("  Node: %s\n", n.Name)

			for k, v := range n.Credentials {
				fmt.Printf("    Credential: %s = %s\n", k, v)
			}
		}
	}
}
