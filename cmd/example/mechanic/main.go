package main

import (
	"github.com/funtimecoding/soil/pkg/web/example/mechanic"
	"github.com/spf13/pflag"
)

func main() {
	address := pflag.String("address", ":8099", "listen address")
	extended := pflag.String("extended", "", "htmx source, empty for default")
	serverSide := pflag.String(
		"server-side",
		"",
		"server-sent event extension source, empty for default",
	)
	pflag.Parse()
	mechanic.Mechanic(*address, *extended, *serverSide)
}
