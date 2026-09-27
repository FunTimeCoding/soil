package constant

import "github.com/funtimecoding/soil/pkg/tool/goclauded/types/timestamp_column"

var TimestampColumns = []*timestamp_column.Column{
	timestamp_column.New(SessionTable, "started_at"),
	timestamp_column.New(SessionTable, "last_seen"),
	timestamp_column.New(SessionTable, "last_active_at"),
	timestamp_column.New(SessionTable, "last_prompt_at"),
	timestamp_column.New(SessionTable, "last_turn_end_at"),
	timestamp_column.New(SessionTable, "closed_at"),
	timestamp_column.New("event", "created_at"),
	timestamp_column.New("queue", "created_at"),
	timestamp_column.New("queue", "consumed_at"),
	timestamp_column.New("notification", "created_at"),
	timestamp_column.New("message", "created_at"),
	timestamp_column.New("completion", "created_at"),
	timestamp_column.New(SummaryTable, "created_at"),
	timestamp_column.New("pulse", "created_at"),
	timestamp_column.New(ContextLoadTable, "occurred_at"),
	timestamp_column.New("tracker_state", "updated_at"),
	timestamp_column.New(RateSnapshotTable, "created_at"),
	timestamp_column.New(RateSnapshotTable, "five_hour_reset"),
	timestamp_column.New(RateSnapshotTable, "seven_day_reset"),
	timestamp_column.New(FableSnapshotTable, "created_at"),
	timestamp_column.New(FableSnapshotTable, "reset_at"),
}
