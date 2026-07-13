package store

import (
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"
	"testing"
	"time"
)

func TestTopAggregates(t *testing.T) {
	s := New(lite.NewMemory())

	if e := s.Create(
		record.Record{
			Fingerprint: "resolved1",
			Name:        "HighMemory",
			Severity:    "critical",
			Start:       time.Now().Add(-10 * time.Minute),
		},
	); e != nil {
		t.Fatalf("create: %v", e)
	}

	if e := s.Resolve("resolved1"); e != nil {
		t.Fatalf("resolve: %v", e)
	}

	if e := s.Create(
		record.Record{
			Fingerprint: "open1",
			Name:        "DiskFull",
			Severity:    "warning",
			Start:       time.Now().Add(-5 * time.Minute),
		},
	); e != nil {
		t.Fatalf("create: %v", e)
	}
	result, e := s.Top(
		0,
		time.Now().Add(-time.Hour),
		time.Now().Add(time.Hour),
	)

	if e != nil {
		t.Fatalf("top: %v", e)
	}

	if len(result) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(result))
	}

	memory := result[0]

	if result[0].Name != "HighMemory" {
		memory = result[1]
	}

	if memory.Count != 1 || memory.CurrentlyFiring != 0 {
		t.Fatalf("unexpected HighMemory aggregate: %+v", memory)
	}

	if memory.AverageDuration < 9*time.Minute ||
		memory.AverageDuration > 11*time.Minute {
		t.Fatalf(
			"average duration outside expected range: %s",
			memory.AverageDuration,
		)
	}

	disk := result[0]

	if result[0].Name != "DiskFull" {
		disk = result[1]
	}

	if disk.CurrentlyFiring != 1 || disk.AverageDuration != 0 {
		t.Fatalf("unexpected DiskFull aggregate: %+v", disk)
	}
}
