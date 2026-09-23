package supervisor

import "fmt"

func (s *Supervisor) handleReloadEnvironment() string {
	if e := s.ReloadEnvironment(); e != nil {
		return fmt.Sprintf("error: %s", e)
	}

	return "ok"
}
