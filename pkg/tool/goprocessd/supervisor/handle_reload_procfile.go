package supervisor

import "fmt"

func (s *Supervisor) handleReloadProcfile() string {
	if e := s.ReloadProcfile(); e != nil {
		return fmt.Sprintf("error: %s", e)
	}

	return "ok"
}
