package supervisor

import (
	"fmt"
	"strings"
)

func (s *Supervisor) handle(line string) string {
	parts := strings.Fields(line)

	if len(parts) == 0 {
		return "error: empty command"
	}

	command := parts[0]
	arguments := parts[1:]

	switch command {
	case "start":
		return s.handleStart(arguments)
	case "stop":
		return s.handleStop(arguments)
	case "restart":
		return s.handleRestart(arguments)
	case "restart-all":
		return s.handleRestartAll()
	case "reload-procfile":
		return s.handleReloadProcfile()
	case "reload-environment":
		return s.handleReloadEnvironment()
	case "log":
		return s.handleLog(arguments)
	case "list":
		return s.handleList()
	case "status":
		return s.handleStatus()
	default:
		return fmt.Sprintf("error: unknown command %s", command)
	}
}
