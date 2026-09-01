package goaudit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/format"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
)

func runTable(
	services []*scan.Service,
	identityWarnings []*concern.Concern,
	clients []*scan.Client,
) {
	fmt.Print(format.Services(services))

	for _, c := range identityWarnings {
		console.Format("%-14s%s\n", c.Path, c.Text)
	}

	console.Line()
	fmt.Print(format.Clients(clients))
}
