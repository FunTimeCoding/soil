package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func timestampColumns() []timestampColumn {
	return []timestampColumn{
		{constant.SessionTable, "started_at"},
		{constant.SessionTable, "last_seen"},
		{constant.SessionTable, "last_active_at"},
		{constant.SessionTable, "closed_at"},
		{"event", "created_at"},
		{"queue", "created_at"},
		{"queue", "consumed_at"},
		{"notification", "created_at"},
		{"message", "created_at"},
		{"completion", "created_at"},
		{constant.SummaryTable, "created_at"},
		{"pulse", "created_at"},
		{constant.ContextLoadTable, "occurred_at"},
		{"tracker_state", "updated_at"},
		{constant.RateSnapshotTable, "created_at"},
		{constant.RateSnapshotTable, "five_hour_reset"},
		{constant.RateSnapshotTable, "seven_day_reset"},
		{constant.FableSnapshotTable, "created_at"},
		{constant.FableSnapshotTable, "reset_at"},
	}
}
