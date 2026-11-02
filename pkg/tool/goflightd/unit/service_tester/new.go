package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/option"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/service"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/unit/mock_system"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	y := mock_system.New()
	t.Cleanup(y.Close)
	m := lite.NewMemory()
	t.Cleanup(func() { relational.CloseMapper(m) })
	migrate.AutoMigrate(m)
	s := store.New(m)
	o := option.New()

	return &Tester{
		Service: service.New(
			s,
			logger.New(t.Context()),
			o.Predicate,
			o.Retention,
			o.Clock,
			y.Run,
			y.Stream,
		),
		Store:    s,
		System:   y,
		Reporter: memory.New(),
		mapper:   m,
	}
}
