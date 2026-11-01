package mock_service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/intent"
)

func (s *Service) Reconcile(
	_ context.Context,
	hypervisor string,
	_ bool,
	_ bool,
) (*intent.Plan, error) {
	return intent.NewPlan(hypervisor), nil
}
