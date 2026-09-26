package goagent

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/client"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

func submit(newClient func() *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "submit <file>",
		Short: "Submit an intent file to the runner",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			content, e := os.ReadFile(arguments[0])
			errors.PanicOnError(e)
			code := newClient().Submit(string(content))

			switch code {
			case http.StatusAccepted:
				fmt.Println("intent submitted")
			case http.StatusConflict:
				fmt.Println("already running")
				os.Exit(1)
			default:
				fmt.Printf("submit failed (%d)\n", code)
				os.Exit(1)
			}
		},
	}
}
