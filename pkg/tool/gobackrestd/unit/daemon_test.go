package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/backrest/face"
	"github.com/funtimecoding/soil/pkg/backrest/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/option"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/unit/daemon_tester"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"go.uber.org/goleak"
	"strings"
	"testing"
)

func TestClockGaugesLatestBackups(t *testing.T) {
	c := mock_client.New()
	c.Add(fixtureStanzas("backed_up.json")...)
	m := metric.New()

	for _, w := range daemon_tester.New(t, c, m, memory.New()).Workers() {
		w.Start()
		w.Close()
	}

	assert.FatalOnError(
		t,
		testutil.GatherAndCompare(
			m.Registry(),
			strings.NewReader(
				`
# HELP pgbackrest_stanza_status Stanza status code from pgbackrest info, 0 is ok
# TYPE pgbackrest_stanza_status gauge
pgbackrest_stanza_status{stanza="fixture"} 0
# HELP pgbackrest_backup_completion_timestamp_seconds Unix time the latest backup of a type finished
# TYPE pgbackrest_backup_completion_timestamp_seconds gauge
pgbackrest_backup_completion_timestamp_seconds{stanza="fixture",type="diff"} 1.794359141e+09
pgbackrest_backup_completion_timestamp_seconds{stanza="fixture",type="full"} 1.794359139e+09
# HELP pgbackrest_backup_size_bytes Database size covered by the latest backup of a type
# TYPE pgbackrest_backup_size_bytes gauge
pgbackrest_backup_size_bytes{stanza="fixture",type="diff"} 2.7360975e+07
pgbackrest_backup_size_bytes{stanza="fixture",type="full"} 2.3690561e+07
# HELP pgbackrest_backup_repository_size_bytes Repository size of the latest backup of a type, with the backups it references
# TYPE pgbackrest_backup_repository_size_bytes gauge
pgbackrest_backup_repository_size_bytes{stanza="fixture",type="diff"} 3.132628e+06
pgbackrest_backup_repository_size_bytes{stanza="fixture",type="full"} 2.824579e+06
# HELP pgbackrest_backup_error 1 when the latest backup of a type found page checksum errors
# TYPE pgbackrest_backup_error gauge
pgbackrest_backup_error{stanza="fixture",type="diff"} 0
pgbackrest_backup_error{stanza="fixture",type="full"} 0
`,
			),
		),
	)
}

func TestClockGaugesStatusOfStanzaWithoutBackups(t *testing.T) {
	c := mock_client.New()
	c.Add(fixtureStanzas("no_backup.json")...)
	m := metric.New()

	for _, w := range daemon_tester.New(t, c, m, memory.New()).Workers() {
		w.Start()
		w.Close()
	}

	assert.FatalOnError(
		t,
		testutil.GatherAndCompare(
			m.Registry(),
			strings.NewReader(
				`
# HELP pgbackrest_stanza_status Stanza status code from pgbackrest info, 0 is ok
# TYPE pgbackrest_stanza_status gauge
pgbackrest_stanza_status{stanza="fixture"} 2
`,
			),
		),
	)
}

func TestDaemonHandsOutGuardFirst(t *testing.T) {
	_, okay := daemon_tester.New(
		t,
		mock_client.New(),
		metric.New(),
		memory.New(),
	).Workers()[0].(*peer.Guard)
	assert.True(t, okay)
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	before := goleak.IgnoreCurrent()
	c := mock_client.New()
	o := option.New()
	o.Source = func() face.BackrestClient { return c }
	d := daemon.New(o, logger.New(t.Context()), memory.New(), metric.New())
	workers := d.Workers()

	for _, w := range workers {
		w.Start()
	}

	for _, w := range workers {
		w.Close()
	}

	d.Close()
	goleak.VerifyNone(t, before)
}
