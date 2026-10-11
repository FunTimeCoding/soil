package gotechnitium

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/technitium"
	"github.com/funtimecoding/soil/pkg/technitium/constant"
	"github.com/funtimecoding/soil/pkg/technitium/policy"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/spf13/cobra"
	"slices"
)

func allowUpdate(
	c *technitium.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var keyName string
	var network string
	var domain string
	var types []string
	result := &cobra.Command{
		Use:   "allow-update [zone]",
		Short: "Allow dynamic updates to a zone from a network with a TSIG key",
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

			if domain == "" {
				domain = join.Empty(constant.WildcardPrefix, arguments[0])
			}

			technitium.AllowUpdate(
				o,
				network,
				policy.New(keyName, domain, types),
			)
			c.MustSetZoneOptions(o)
			printZoneOptions(c.MustZoneOptions(arguments[0]))
		},
	}
	f := result.Flags()
	f.StringVar(&keyName, argumentConstant.Key, "", "TSIG key name")
	f.StringVar(
		&network,
		argumentConstant.Network,
		"",
		"Network allowed to send updates, such as 198.51.100.0/24",
	)
	f.StringVar(
		&domain,
		argumentConstant.Domain,
		"",
		"Domain the key may update; default every name below the zone",
	)
	f.StringSliceVar(
		&types,
		argumentConstant.Type,
		nil,
		"Record types the key may update, comma separated",
	)
	errors.PanicOnError(result.MarkFlagRequired(argumentConstant.Key))
	errors.PanicOnError(result.MarkFlagRequired(argumentConstant.Network))
	errors.PanicOnError(result.MarkFlagRequired(argumentConstant.Type))

	return result
}
