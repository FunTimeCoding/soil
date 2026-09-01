package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/gitlab"
)

func Feature() {
	g := gitlab.NewEnvironment()

	for _, f := range g.MustFeatures() {
		console.Format("Feature: %+v\n", f)
	}

	for _, d := range g.MustFeatureDefinitions() {
		console.Format("Definition: %+v\n", d)
	}
}
