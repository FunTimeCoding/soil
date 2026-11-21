package upgrade_host

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"

func New(
	s *system_result.Summary,
	f *system_result.Finding,
) *Host {
	return &Host{
		Name:      s.Name,
		Priority:  s.Priority,
		Labels:    s.Labels,
		Installed: f.Installed,
		Candidate: f.Candidate,
	}
}
