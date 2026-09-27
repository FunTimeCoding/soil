package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/collector/bluetooth"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/collector/wireless"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/event"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/mark"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/snapshot"
	"testing"
	"time"
)

func TestBluetoothParse(t *testing.T) {
	result := bluetooth.Parse(constant.BluetoothSample)
	assert.Integer(t, 2, len(result))
	assert.String(t, "connected", result["Example Mouse"])
	assert.String(t, "disconnected", result["Example Keyboard"])
}

func TestBluetoothParseInvalid(t *testing.T) {
	assert.Integer(t, 0, len(bluetooth.Parse("not json")))
}

func TestWirelessParse(t *testing.T) {
	result := wireless.Parse(constant.WirelessSample)
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
	result := wireless.Parse("Total Time : ignored\n")
	assert.Integer(t, 0, len(result))
}

func TestStoreRoundTrip(t *testing.T) {
	s := store.New(lite.NewMemory())
	defer s.Close()
	now := time.Now()
	s.MustCreateEvent(
		event.Event{
			Time:      now,
			Process:   "ensembled",
			Subsystem: "com.apple.ensemble",
			Message:   "session started",
		},
	)
	s.MustCreateEvent(
		event.Event{
			Time:    now.Add(-2 * time.Hour),
			Process: "rapportd",
			Message: "old event",
		},
	)
	s.MustCreateSnapshot(
		snapshot.Snapshot{
			Time:  now,
			Kind:  "wireless",
			Key:   "awdl.Channel Sequence",
			Value: "{(44, 80MHz)}",
		},
	)
	events, e := s.EventsByTimeRange(
		now.Add(-1*time.Hour),
		now.Add(time.Hour),
		100,
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 1, len(events))
	assert.String(t, "ensembled", events[0].Process)
	snapshots, f := s.SnapshotsByTimeRange(
		now.Add(-1*time.Hour),
		now.Add(time.Hour),
		100,
	)
	assert.FatalOnError(t, f)
	assert.Integer(t, 1, len(snapshots))
	assert.String(t, "wireless", snapshots[0].Kind)
	last, g := s.LastEventTime()
	assert.FatalOnError(t, g)
	assert.NotNil(t, last)
}

func TestStoreMarks(t *testing.T) {
	s := store.New(lite.NewMemory())
	defer s.Close()
	first := mark.New(time.Now().Add(-time.Minute), "broke", "")
	assert.FatalOnError(t, s.CreateMark(first))
	second := mark.New(time.Now(), "recovered", "mouse cycled")
	assert.FatalOnError(t, s.CreateMark(second))
	marks, e := s.RecentMarks(10)
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, len(marks))
	assert.String(t, "recovered", marks[0].Label)
	count, f := s.MarkCount()
	assert.FatalOnError(t, f)
	assert.Integer(t, 2, count)
}
