package unit

import (
	"bufio"
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/unit/base"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"testing"
	"time"
)

func TestLiveDashboardHearsWrite(t *testing.T) {
	s := base.New(t)
	x, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()
	q, e := http.NewRequestWithContext(
		x,
		http.MethodGet,
		s.Address(
			join.Empty(
				webConstant.LivePath,
				"?",
				webConstant.ParameterSubscribe,
				"=",
				constant.EventRecent,
			),
		),
		nil,
	)
	assert.FatalOnError(t, e)
	r, f := s.Client().Do(q)
	assert.FatalOnError(t, f)
	defer errors.PanicClose(r.Body)
	stream := bufio.NewReader(r.Body)
	awaitLine(t, stream, join.Empty("event: ", constant.EventRecent))
	_, g := s.MaintenanceClient(t).PostEntry(
		x,
		client.PostEntryRequest{Action: "restart", User: "zulu"},
	)
	assert.FatalOnError(t, g)
	awaitLine(t, stream, "zulu")
}
