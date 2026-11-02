package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/salt"
	"github.com/funtimecoding/soil/pkg/provision/salt/salt_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/option"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/service"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"path/filepath"
	"testing"
	"time"
)

func TestKeysBeforeConnectingAnswersNotConnected(t *testing.T) {
	_, e := service_tester.New(t).Keys()
	assert.True(t, errors.Is(e, constant.ErrorNotConnected))
}

func TestKeysAfterLoginAnswerTheMastersKeys(t *testing.T) {
	m := salt_tester.New(t, "bravo")
	host, port := upstream_tester.HostPort(t, m.URL)
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "gosaltd.db")
	o.Master = func() *salt.Client {
		return salt.New(host, port, "alfa", "bravo", salt.WithInsecure())
	}
	s := service.New(o, logger.New(t.Context()), memory.New())
	t.Cleanup(s.Close)

	for _, w := range s.Workers() {
		w.Start()
		t.Cleanup(w.Close)
	}

	deadline := time.Now().Add(5 * time.Second)

	for time.Now().Before(deadline) {
		keys, e := s.Keys()

		if e == nil {
			assert.Strings(t, []string{"alfa"}, keys.Minions)

			return
		}

		time.Sleep(20 * time.Millisecond)
	}

	t.Fatal("keys never answered after login")
}

func TestServiceReadsRunsFromItsOwnDatabase(t *testing.T) {
	v, e := service_tester.New(t).Recent(10)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, v)
}

func TestServiceClosesAStartedRunner(t *testing.T) {
	s := service_tester.New(t)

	for _, w := range s.Workers() {
		w.Start()
		t.Cleanup(w.Close)
	}
}
