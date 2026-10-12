package gotechnitium

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/technitium"
	technitiumConstant "github.com/funtimecoding/soil/pkg/technitium/constant"
	"github.com/funtimecoding/soil/pkg/technitium/key"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/spf13/cobra"
)

func addKey(
	c *technitium.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var algorithm string
	result := &cobra.Command{
		Use:   "add-key [name]",
		Short: "Generate a TSIG key on the server and print only its secret",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			keys := c.MustKeys()

			if key.Find(keys, arguments[0]) != nil {
				t.Exitf(
					"key %s exists, its secret is not replaced\n",
					arguments[0],
				)

				return
			}

			k := key.Generate(arguments[0], algorithm)
			c.MustSetKeys(append(keys, k))
			console.Line(k.Secret)
		},
	}
	result.Flags().StringVar(
		&algorithm,
		argumentConstant.Algorithm,
		technitiumConstant.DefaultAlgorithm,
		"TSIG algorithm",
	)

	return result
}
