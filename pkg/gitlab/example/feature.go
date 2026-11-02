package example

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/gitlab"
)

func Feature() {
	g := gitlab.NewEnvironment()
	defer g.Close()
	x := context.Background()

	for _, f := range g.MustFeatures(x) {
		console.Format("Feature: %+v\n", f)
	}

	for _, d := range g.MustFeatureDefinitions(x) {
		console.Format("Definition: %+v\n", d)
	}
}
