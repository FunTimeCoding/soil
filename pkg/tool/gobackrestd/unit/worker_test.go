package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/types/summary"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/worker"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"strings"
	"testing"
	"time"
)

func TestPollDropsStanzaNoLongerListed(t *testing.T) {
	stanzas := []string{"alfa", "bravo"}
	y := prometheus.NewRegistry()
	w := worker.New(
		func(context.Context) ([]*summary.Summary, error) {
			s := summary.New(stanzas[0], 0)
			s.Latest["full"] = fixtureStanzas("backed_up.json")[0].Backup[0]
			stanzas = stanzas[1:]

			return []*summary.Summary{s}, nil
		},
		time.Minute,
		logger.New(t.Context()),
		y,
		memory.New(),
	)
	defer w.Close()
	w.Poll()
	w.Poll()
	assert.FatalOnError(
		t,
		testutil.GatherAndCompare(
			y,
			strings.NewReader(
				`
# HELP pgbackrest_stanza_status Stanza status code from pgbackrest info, 0 is ok
# TYPE pgbackrest_stanza_status gauge
pgbackrest_stanza_status{stanza="bravo"} 0
# HELP pgbackrest_backup_completion_timestamp_seconds Unix time the latest backup of a type finished
# TYPE pgbackrest_backup_completion_timestamp_seconds gauge
pgbackrest_backup_completion_timestamp_seconds{stanza="bravo",type="full"} 1.794359139e+09
`,
			),
			constant.StatusMetric,
			constant.CompletionMetric,
		),
	)
}
