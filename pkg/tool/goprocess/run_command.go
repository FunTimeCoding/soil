package goprocess

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/client"
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
  log PROCESS            Show the current generation of log output
  log PROCESS all        Show all retained log output
  log PROCESS clear      Discard retained log output
  list                   List process names
  status                 Show running status
`,
		),
		Args: cobra.MinimumNArgs(1),
		RunE: func(
			m *cobra.Command,
			arguments []string,
		) error {
			m.SilenceUsage = true
			procfilePath, e := m.Flags().GetString("file")
			errors.PanicOnError(e)
			c := client.NewLocal(procfilePath)
			defer c.Close()

			return Dispatch(c, arguments[0], arguments[1:])
		},
	}
}
