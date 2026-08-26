package server

import "fmt"

func (s *Server) handleLog(arguments []string) string {
	if len(arguments) == 0 {
		return "error: log requires a process name"
	}

	name := arguments[0]
	p := s.findProcess(name)

	if p == nil {
		return fmt.Sprintf("error: unknown process %s", name)
	}

	if len(arguments) == 1 {
		return currentLog(p, name)
	}

	switch arguments[1] {
	case "all":
		return lines(p.Log())
	case "clear":
		p.ClearLog()

		return "ok"
	default:
		return fmt.Sprintf("error: unknown log option %s", arguments[1])
	}
}
