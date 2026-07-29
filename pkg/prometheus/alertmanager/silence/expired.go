package silence

import "github.com/funtimecoding/soil/pkg/prometheus/constant"

func (s *Silence) Expired() bool {
	return s.State == constant.ExpiredState
}
