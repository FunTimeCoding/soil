package affected_row

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"

func New(
	s *system_result.Summary,
	f *system_result.Finding,
	kind string,
) *Row {
	return &Row{
		System:    s.Name,
		Coverage:  s.Coverage,
		Priority:  s.Priority,
		Labels:    s.Labels,
		Source:    f.Source,
		Installed: f.Installed,
		Fixed:     f.Fixed,
		Candidate: f.Candidate,
		Kind:      kind,
	}
}
