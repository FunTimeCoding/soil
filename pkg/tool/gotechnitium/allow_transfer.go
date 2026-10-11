package gotechnitium

import (
	"github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/technitium"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/spf13/cobra"
	"slices"
)

func allowTransfer(
	c *technitium.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var keyName string
	var network string
	result := &cobra.Command{
		Use:   "allow-transfer [zone]",
		Short: "Allow zone transfers to a network with a TSIG key",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			o := c.MustZoneOptions(arguments[0])

			if !slices.Contains(o.AvailableKeyNames, keyName) {
				t.Exitf("no key %s on the server, add it first\n", keyName)

				return
			}

			technitium.AllowTransfer(o, network, keyName)
			c.MustSetZoneOptions(o)
			printZoneOptions(c.MustZoneOptions(arguments[0]))
		},
	}
	f := result.Flags()
	f.StringVar(&keyName, constant.Key, "", "TSIG key name")
	f.StringVar(
		&network,
		constant.Network,
		"",
		"Network allowed to transfer the zone, such as 198.51.100.0/24",
	)
	errors.PanicOnError(result.MarkFlagRequired(constant.Key))
	errors.PanicOnError(result.MarkFlagRequired(constant.Network))

	return result
}
