package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	generated "github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/spf13/cobra"
)

func updateConfigurationContext(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var payload string
	var weight int32
	result := &cobra.Command{
		Use:   "update-configuration-context [name]",
		Short: "Update a configuration context; a given payload or assignment replaces the old one whole",
		Args:  cobra.ExactArgs(1),
	}
	assignment := assignmentFlags(result)
	result.Run = func(
		o *cobra.Command,
		arguments []string,
	) {
		body := generated.UpdateConfigurationContextRequest{
			Assignment: assignment(),
		}

		if o.Flags().Changed("payload") {
			p := map[string]any{}
			errors.PanicOnError(notation.Decode(payload, &p))
			body.Payload = &p
		}

		if o.Flags().Changed("weight") {
			body.Weight = &weight
		}

		t.Emit(c.UpdateConfigurationContext(arguments[0], body))
	}
	result.Flags().StringVar(
		&payload,
		"payload",
		"",
		"JSON object replacing the whole payload",
	)
	result.Flags().Int32Var(
		&weight,
		"weight",
		0,
		"merge order, higher weights win",
	)

	return result
}
