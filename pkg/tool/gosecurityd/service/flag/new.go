package flag

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/vulnerability"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store"
)

func New(
	s *store.Store,
	v *vulnerability.Service,
) *Service {
	return &Service{store: s, vulnerability: v}
}
