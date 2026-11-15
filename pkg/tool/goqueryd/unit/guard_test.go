package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"testing"
)

func TestGuard(t *testing.T) {
	w := base.Serve(
		t,
		service_tester.NewOption(
			t,
			service_tester.NewEmbedder(t),
			service_tester.NewReranker(t),
		),
		memory.New(),
	)
	w.VerifyBase(t)
	w.VerifyFavicon(t)
	w.VerifyGuarded(t, "/api/status")
	w.VerifyOpen(t, constant.DashboardPath)
	w.VerifyModelContext(t)
}
