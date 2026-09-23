package cross_service_tester

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/connector"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
	"time"
)

func New(t *testing.T) *Tester {
	t.Helper()
	coordinator := base.New(t)
	now := time.Now().UTC()

	return &Tester{
		t: t,
		Service: service.New(
			store.New(lite.NewMemory(), func() time.Time { return now }),
			notifier.New(),
		),
		Coordinator: connector.New(
			locator.New(
				constant.Localhost,
			).Insecure().Port(coordinator.Port).String(),
			true,
			generative.ModelContextTestToken,
		),
		Goclauded: coordinator,
		now:       &now,
	}
}
