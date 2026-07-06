package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/gitlab"
)

func Feature() {
	g := gitlab.NewEnvironment()

	for _, f := range g.MustFeatures() {
		fmt.Printf("Feature: %+v\n", f)
	}

	for _, d := range g.MustFeatureDefinitions() {
		fmt.Printf("Definition: %+v\n", d)
	}
}
