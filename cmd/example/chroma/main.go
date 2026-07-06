//go:build local

package main

import (
	"fmt"
	"github.com/amikos-tech/chroma-go/pkg/api/v2"
	"github.com/funtimecoding/soil/pkg/generative/chroma"
	"github.com/funtimecoding/soil/pkg/generative/chroma/example"
)

func main() {
	// Cannot make cmd/gochroma package because it would require CGO_ENABLED=1
	//  Solve CGO requirement in chroma library or create own chroma project which uses CGO_ENABLED=1
	c := chroma.NewEnvironment()
	defer c.Close()

	for _, d := range c.Databases(v2.NewDefaultTenant()) {
		fmt.Printf("Database: %s\n", d.Name())
	}

	for _, o := range c.Collections() {
		fmt.Printf("Collection: %s\n", o.Name())
	}

	if false {
		example.Collection()
		example.Client()
	}
}
