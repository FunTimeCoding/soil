package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/service"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/event"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/snapshot"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/unit/mock_system"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/worker"
	"net/http"
	"testing"
	"time"
)

func TestTheServiceClosesItsDatabaseOnce(t *testing.T) {
	v := service.New(
		base.NewOption(t, mock_system.New()),
		logger.New(t.Context()),
	)
	v.Close()
	v.Close()
	_, e := v.Status()
	assert.Error(t, e)
}

func TestTheWorkerClosesTwice(t *testing.T) {
	s := base.New(t)
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

func TestABluetoothSampleRecordsOnlyChanges(t *testing.T) {
	s := base.New(t)
	s.System.Answer(constant.SystemProfiler, constant.BluetoothSample)
	assert.Count(t, 0, s.Service.SampleBluetooth())
	assert.Count(t, 0, s.Service.SampleBluetooth())
	v, e := s.Service.Status()
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, v.Snapshots)
}

func TestAWirelessSampleRecordsOnlyChanges(t *testing.T) {
	s := base.New(t)
	s.System.Answer(constant.Sudo, constant.WirelessSample)
	assert.True(t, s.Service.WirelessAvailable())
	assert.Count(t, 0, s.Service.SampleWireless())
	assert.Count(t, 0, s.Service.SampleWireless())
	v, e := s.Service.Status()
	assert.FatalOnError(t, e)
	assert.Integer(t, 6, v.Snapshots)
}

func TestAFailedCommandIsLoggedNotReported(t *testing.T) {
	s := base.New(t)
	s.System.Fail(constant.SystemProfiler, errors.Format("exit", 1))
	s.System.Fail(constant.Sudo, errors.Format("exit", 1))
	assert.False(t, s.Service.WirelessAvailable())
	worker.NewEvery(
		s.Service.SampleBluetooth,
		logger.New(t.Context()),
		s.Reporter,
		time.Hour,
	).Poll()
	assert.Count(t, 0, s.Service.SampleWireless())
	assert.Count(t, 0, s.Reporter.Events())
}

func TestAFailingStoreIsReportedByTheWorker(t *testing.T) {
	s := base.New(t)
	s.System.Answer(constant.SystemProfiler, constant.BluetoothSample)
	s.Service.Close()
	worker.NewEvery(
		s.Service.SampleBluetooth,
		logger.New(t.Context()),
		s.Reporter,
		time.Hour,
	).Poll()
	assert.Count(t, 1, s.Reporter.Events())
}

func TestPruneDropsWhatIsPastRetention(t *testing.T) {
	s := base.New(t)
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

func TestAFailingStoreIsAnsweredAndReported(t *testing.T) {
	s := base.New(t)
	s.Service.Close()
	r, e := s.Client.GetStatusWithResponse(context.Background())
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusInternalServerError, r.StatusCode())
	assert.NotNil(t, r.JSON500)
	assert.Count(t, 1, s.Reporter.Events())
}

func TestAMarkWithoutLabelIsRefused(t *testing.T) {
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
