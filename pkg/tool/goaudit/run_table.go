package goaudit

import (
	"fmt"
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
		fmt.Printf("%-14s%s\n", c.Path, c.Text)
	}

	fmt.Println()
	fmt.Print(format.Clients(clients))
}
