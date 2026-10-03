//go:build local

package main

import (
	"github.com/amikos-tech/chroma-go/pkg/api/v2"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/generative/chroma"
	"github.com/funtimecoding/soil/pkg/generative/chroma/example"
)

func main() {
	c := chroma.NewEnvironment()
	defer c.Close()

	for _, d := range c.Databases(v2.NewDefaultTenant()) {
		console.Format("Database: %s\n", d.Name())
	}

	for _, o := range c.Collections() {
		console.Format("Collection: %s\n", o.Name())
	}

	if false {
		example.Collection()
		example.Client()
	}
}
