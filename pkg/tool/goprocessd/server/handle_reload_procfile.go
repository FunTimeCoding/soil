package server

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/process"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/procfile"
)

func (s *Server) handleReloadProcfile() string {
	entries, e := procfile.Parse(s.procfilePath)

	if e != nil {
		return fmt.Sprintf("error: %s", e)
	}

	current := make(map[string]*process.Process, len(s.processes))

	for _, p := range s.processes {
		current[p.Name] = p
	}

	wanted := make(map[string]procfile.Entry, len(entries))

	for _, entry := range entries {
		wanted[entry.Name] = entry

		if len(entry.Name) > s.maxNameWidth {
			s.maxNameWidth = len(entry.Name)
		}
	}

	var result []*process.Process

	for _, p := range s.processes {
		entry, exists := wanted[p.Name]

		if !exists {
			if e := p.Stop(); e != nil {
				return fmt.Sprintf("error: %s", e)
			}

			continue
		}

		if p.Command != entry.Command {
			if e := p.Stop(); e != nil {
				return fmt.Sprintf("error: %s", e)
			}

			replacement := process.New(
				entry.Name,
				entry.Command,
				p.ColorIndex,
				s.maxNameWidth,
			)
			replacement.Spawn(s.environment.Build(), nil)
			result = append(result, replacement)

			continue
		}

		result = append(result, p)
	}

	for _, entry := range entries {
		if current[entry.Name] != nil {
			continue
		}

		p := process.New(
			entry.Name,
			entry.Command,
			len(result),
			s.maxNameWidth,
		)
		p.Spawn(s.environment.Build(), nil)
		result = append(result, p)
	}

	s.processes = result

	return "ok"
}
