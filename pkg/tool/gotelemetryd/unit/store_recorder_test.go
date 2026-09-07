package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/telemetry/record"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store_recorder"
	"testing"
)

func TestStoreRecorder(t *testing.T) {
	s := store.New(lite.NewMemory())
	r := store_recorder.New(s, memory.New())
	e := record.NewBaseline(
		"GetSummary",
		constant.SurfaceWebService,
		"tester",
		constant.OutcomeSuccess,
	)
	e.Detail = map[string]string{"operation": "GetSummary"}
	r.Record(e)
	events, queryError := s.Recent(store.NewQueryOption())

	if queryError != nil {
		t.Fatalf("recent: %v", queryError)
	}

	if len(events) != 1 {
		t.Fatalf("expected one event, got %d", len(events))
	}

	if events[0].Tool != "GetSummary" {
		t.Fatalf("unexpected tool: %s", events[0].Tool)
	}

	if events[0].Detail == nil {
		t.Fatal("expected detail to be set")
	}
}
