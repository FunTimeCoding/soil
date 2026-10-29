package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/parse"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/event"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/mark"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/snapshot"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/unit/base"
	"testing"
	"time"
)

func TestBluetoothParse(t *testing.T) {
	result := parse.Bluetooth(constant.BluetoothSample)
	assert.Integer(t, 2, len(result))
	assert.String(t, "connected", result["Example Mouse"])
	assert.String(t, "disconnected", result["Example Keyboard"])
}

func TestBluetoothParseInvalid(t *testing.T) {
	assert.Integer(t, 0, len(parse.Bluetooth("not json")))
}

func TestWirelessParse(t *testing.T) {
	result := parse.Wireless(constant.WirelessSample)
	assert.Integer(t, 6, len(result))
	assert.String(t, "en0 (Wi-Fi)", result["network.Primary IPv4"])
	assert.String(t, "aa:bb:cc:dd:ee:ff", result["wifi.MAC Address"])
	assert.String(t, "-45 dBm", result["wifi.RSSI"])
	assert.String(t, "5g44/80", result["wifi.Channel"])
	assert.String(t, "Yes", result["awdl.AWDL Enabled"])
	assert.String(
		t,
		"{(44, 80MHz), (6, 20MHz)}",
		result["awdl.Channel Sequence"],
	)
}

func TestWirelessParseSkipsPreamble(t *testing.T) {
	result := parse.Wireless("Total Time : ignored\n")
	assert.Integer(t, 0, len(result))
}

func TestStoreRoundTrip(t *testing.T) {
	s := base.New(t)
	now := time.Now()
	assert.FatalOnError(
		t,
		s.Store.CreateEvent(
			event.Event{
				Time:      now,
				Process:   "ensembled",
				Subsystem: "com.apple.ensemble",
				Message:   "session started",
			},
		),
	)
	assert.FatalOnError(
		t,
		s.Store.CreateEvent(
			event.Event{
				Time:    now.Add(-2 * time.Hour),
				Process: "rapportd",
				Message: "old event",
			},
		),
	)
	assert.FatalOnError(
		t,
		s.Store.CreateSnapshot(
			snapshot.Snapshot{
				Time:  now,
				Kind:  "wireless",
				Key:   "awdl.Channel Sequence",
				Value: "{(44, 80MHz)}",
			},
		),
	)
	events, e := s.Service.Events(
		now.Add(-1*time.Hour),
		now.Add(time.Hour),
		100,
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 1, len(events))
	assert.String(t, "ensembled", events[0].Process)
	snapshots, f := s.Service.Snapshots(
		now.Add(-1*time.Hour),
		now.Add(time.Hour),
		100,
	)
	assert.FatalOnError(t, f)
	assert.Integer(t, 1, len(snapshots))
	assert.String(t, "wireless", snapshots[0].Kind)
	v, g := s.Service.Status()
	assert.FatalOnError(t, g)
	assert.NotNil(t, v.LastEvent)
}

func TestStoreMarks(t *testing.T) {
	s := base.New(t)
	assert.FatalOnError(
		t,
		s.Store.CreateMark(mark.New(time.Now().Add(-time.Minute), "broke", "")),
	)
	_, f := s.Service.CreateMark("recovered", "mouse cycled")
	assert.FatalOnError(t, f)
	marks, g := s.Service.Marks(10)
	assert.FatalOnError(t, g)
	assert.Integer(t, 2, len(marks))
	assert.String(t, "recovered", marks[0].Label)
	v, h := s.Service.Status()
	assert.FatalOnError(t, h)
	assert.Integer(t, 2, v.Marks)
}
