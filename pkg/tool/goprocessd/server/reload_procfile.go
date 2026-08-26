package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/process"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/procfile"
)

func (s *Server) ReloadProcfile() error {
	s.commandMutex.Lock()
	defer s.commandMutex.Unlock()
	entries, e := procfile.Parse(s.procfilePath)

	if e != nil {
		return e
	}

	existing := s.snapshotProcesses()
	current := make(map[string]*process.Process, len(existing))

	for _, p := range existing {
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

	for _, p := range existing {
		entry, exists := wanted[p.Name]

		if !exists {
			if f := p.Stop(); f != nil {
				return f
			}

			continue
		}

		if p.Command != entry.Command {
			if f := p.Stop(); f != nil {
				return f
			}

			replacement := process.Replace(p, entry.Command)
			s.spawn(replacement)
			result = append(result, replacement)

			continue
		}

		result = append(result, p)
	}

	for _, entry := range entries {
		if current[entry.Name] != nil {
			continue
		}

		p := process.New(entry.Name, entry.Command, len(result), s.maxNameWidth)
		s.spawn(p)
		result = append(result, p)
	}

	s.setProcesses(result)

	return nil
}
