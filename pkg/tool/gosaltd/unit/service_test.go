package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/provision/salt"
	"github.com/funtimecoding/soil/pkg/provision/salt/face"
	"github.com/funtimecoding/soil/pkg/provision/salt/salt_server"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"testing"
	"time"
)

func TestKeysWhileMasterIsDownAnswerUnavailable(t *testing.T) {
	v, _ := service_tester.New(t, service_tester.NewOption(t))
	_, e := v.Keys()
	assert.True(t, unavailable.Is(e))
}

func TestKeysAfterLoginAnswerMastersKeys(t *testing.T) {
	m := salt_server.New(t, "bravo")
	host, port := upstream_server.HostPort(t, m.URL)
	o := service_tester.NewOption(t)
	o.Master = func() face.SaltClient {
		return salt.New(host, port, "alfa", "bravo", salt.WithInsecure())
	}
	v, r := service_tester.New(t, o)
	r.Start()
	deadline := time.Now().Add(5 * time.Second)

	for time.Now().Before(deadline) {
		keys, e := v.Keys()

		if e == nil {
			assert.Strings(t, []string{"alfa"}, keys.Minions)

			return
		}

		time.Sleep(20 * time.Millisecond)
	}

	t.Fatal("keys never answered after login")
}

func TestServiceReadsRunsFromItsOwnDatabase(t *testing.T) {
	v, _ := service_tester.New(t, service_tester.NewOption(t))
	runs, e := v.Recent(10)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, runs)
}
