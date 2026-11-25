package snapshot

import gitlabConstant "github.com/funtimecoding/soil/pkg/gitlab/constant"

func (s *Snapshot) Finished() bool {
	return s.State == gitlabConstant.MergedState ||
		s.State == gitlabConstant.ClosedState
}
