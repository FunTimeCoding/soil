package goprocess

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/client"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/socket"
	"github.com/spf13/cobra"
	"strings"
)

func runCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "run COMMAND [ARGUMENTS...]",
		Short: "Send a command to the running server",
		Long: strings.TrimSpace(
			`
Send a command to the running goprocessd server.

Commands:
  start PROCESS...       Start named processes
  stop PROCESS...        Stop named processes
  restart PROCESS...     Restart named processes
  restart-all            Restart all processes
  reload-procfile        Re-read Procfile, start/stop/restart as needed
  reload-environment     Re-eval .envrc for future restarts
  log PROCESS            Show recent log output
  list                   List process names
  status                 Show running status
`,
		),
		Args: cobra.MinimumNArgs(1),
		RunE: func(
			cmd *cobra.Command,
			arguments []string,
		) error {
			cmd.SilenceUsage = true
			procfilePath, e := cmd.Flags().GetString("file")
			errors.PanicOnError(e)
			command := arguments[0]
			commandArguments := arguments[1:]
			response, e := client.Send(
				socket.Path(procfilePath),
				command,
				commandArguments,
			)

			if e != nil {
				return e
			}

			if strings.HasPrefix(response, "error:") {
				return fmt.Errorf(
					"%s",
					strings.TrimPrefix(response, "error: "),
				)
			}

			if response != "ok" {
				fmt.Println(response)
			}

			return nil
		},
	}
}
