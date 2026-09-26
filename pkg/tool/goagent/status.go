package goagent

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/client"
	"github.com/spf13/cobra"
)

func status(newClient func() *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Check runner state",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			s, e := newClient().Status()

			if unreachable.Is(e) {
				fmt.Println("runner unreachable")

				return
			}

			errors.PanicOnError(e)

			if s == nil {
				fmt.Println("no response from runner")

				return
			}

			fmt.Printf("state: %s\n", s.State)

			if s.Result != nil && *s.Result != "" {
				fmt.Printf("\n%s\n", *s.Result)
			}
		},
	}
}
