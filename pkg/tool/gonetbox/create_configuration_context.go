package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/spf13/cobra"
)

func createConfigurationContext(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var payload string
	var weight int32
	result := &cobra.Command{
		Use:   "create-configuration-context [name]",
		Short: "Create a configuration context; no assignment applies it to everything",
		Args:  cobra.ExactArgs(1),
	}
	assignment := assignmentFlags(result)
	result.Run = func(
		_ *cobra.Command,
		arguments []string,
	) {
		body := client.CreateConfigurationContextRequest{
			Name:       arguments[0],
			Assignment: assignment(),
		}
		errors.PanicOnError(notation.Decode(payload, &body.Payload))

		if weight != 0 {
			body.Weight = &weight
		}

		t.EmitValue(c.CreateConfigurationContext(context.Background(), body))
	}
	result.Flags().StringVar(
		&payload,
		"payload",
		"{}",
		"JSON object the context contributes",
	)
	result.Flags().Int32Var(
		&weight,
		"weight",
		0,
		"merge order, higher weights win (default 1000)",
	)

	return result
}
