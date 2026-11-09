package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/event"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/snapshot"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/worker"
	"net/http"
	"testing"
	"time"
)

func TestWorkerClosesTwice(t *testing.T) {
	s := service_tester.New(t)
	w := worker.NewEvery(
		s.Service.Prune,
		logger.New(t.Context()),
		s.Reporter,
		time.Hour,
	)
	w.Start()
	w.Close()
	w.Close()
}

func TestBluetoothSampleRecordsOnlyChanges(t *testing.T) {
	s := service_tester.New(t)
	s.System.Answer(constant.SystemProfiler, constant.BluetoothSample)
	assert.Count(t, 0, s.Service.SampleBluetooth())
	assert.Count(t, 0, s.Service.SampleBluetooth())
	v, e := s.Service.Status()
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, v.Snapshots)
}

func TestWirelessSampleRecordsOnlyChanges(t *testing.T) {
	s := service_tester.New(t)
	s.System.Answer(constant.Sudo, constant.WirelessSample)
	assert.True(t, s.Service.WirelessAvailable())
	assert.Count(t, 0, s.Service.SampleWireless())
	assert.Count(t, 0, s.Service.SampleWireless())
	v, e := s.Service.Status()
	assert.FatalOnError(t, e)
	assert.Integer(t, 6, v.Snapshots)
}

func TestFailedCommandIsLoggedNotReported(t *testing.T) {
	s := service_tester.New(t)
	s.System.Fail(constant.SystemProfiler, errors.Format("exit", 1))
	s.System.Fail(constant.Sudo, errors.Format("exit", 1))
	assert.False(t, s.Service.WirelessAvailable())
	owned := worker.NewEvery(
		s.Service.SampleBluetooth,
		logger.New(t.Context()),
		s.Reporter,
		time.Hour,
	)
	defer owned.Close()
	owned.Poll()
	assert.Count(t, 0, s.Service.SampleWireless())
	assert.Count(t, 0, s.Reporter.Events())
}

func TestFailingStoreIsReportedByWorker(t *testing.T) {
	s := service_tester.New(t)
	s.System.Answer(constant.SystemProfiler, constant.BluetoothSample)
	s.CloseStore()
	owned := worker.NewEvery(
		s.Service.SampleBluetooth,
		logger.New(t.Context()),
		s.Reporter,
		time.Hour,
	)
	defer owned.Close()
	owned.Poll()
	assert.Count(t, 1, s.Reporter.Events())
}

func TestPruneDropsWhatIsPastRetention(t *testing.T) {
	s := service_tester.New(t)
	old := time.Now().Add(-constant.Retention - time.Hour)
	assert.FatalOnError(
		t,
		s.Store.CreateEvent(event.Event{Time: old, Process: "rapportd"}),
	)
	assert.FatalOnError(
		t,
		s.Store.CreateEvent(event.Event{Time: time.Now(), Process: "sharingd"}),
	)
	assert.FatalOnError(
		t,
		s.Store.CreateSnapshot(snapshot.Snapshot{Time: old, Kind: "wireless"}),
	)
	assert.Count(t, 0, s.Service.Prune())
	v, e := s.Service.Status()
	assert.FatalOnError(t, e)
	assert.Integer(t, 1, v.Events)
	assert.Integer(t, 0, v.Snapshots)
}

func TestFailingStoreIsAnsweredAndReported(t *testing.T) {
	s := base.New(t)
	assert.FatalOnError(t, s.Mapper.Migrator().DropTable(event.Stub()))
	r, e := s.Client.GetStatusWithResponse(context.Background())
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusInternalServerError, r.StatusCode())
	assert.NotNil(t, r.JSON500)
	assert.Count(t, 1, s.Reporter.Events())
}

func TestMarkWithoutLabelIsRefused(t *testing.T) {
	s := base.New(t)
	r, e := s.Client.PostMarkWithResponse(
		context.Background(),
		client.PostMarkJSONRequestBody{},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusBadRequest, r.StatusCode())
	assert.String(t, "label is required", r.JSON400.Error)
	assert.Count(t, 0, s.Reporter.Events())
}
