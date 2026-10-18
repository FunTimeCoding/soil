package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
)

func newService(
	s *store.Store,
	m face.ReachableEmbedder,
	a queryd.Reranker,
) *service.Service {
	return service.New(
		s,
		m,
		a,
		constant.FixtureEmbedTarget,
		logger.New(context.Background()),
		memory.New(),
	)
}
