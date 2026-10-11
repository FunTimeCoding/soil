package gotechnitium

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/technitium"
	"github.com/spf13/cobra"
)

func listKeys(c *technitium.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-keys",
		Short: "List TSIG keys and their algorithms, never their secrets",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			for _, k := range c.MustKeys() {
				console.Format("%s %s\n", k.Name, k.Algorithm)
			}
		},
	}
}
