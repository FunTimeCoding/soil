package cross_service_tester

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/connector"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
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
				webConstant.Localhost,
			).Insecure().Port(coordinator.Port).String(),
			true,
			generativeConstant.ModelContextTestToken,
		),
		Goclauded: coordinator,
		now:       &now,
	}
}
