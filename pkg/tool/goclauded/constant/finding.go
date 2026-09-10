package constant

const (
	PoolExhausted          = "pool_exhausted"
	MissingTranscriptEmpty = "missing_transcript_empty"
	MissingTranscriptKept  = "missing_transcript_kept"
	OrphanTrackerState     = "orphan_tracker_state"
	UnownedQueue           = "unowned_queue"
	UnownedNotification    = "unowned_notification"
	StaleCallsign          = "stale_callsign"
	MigrationIncomplete    = "migration_incomplete"
)

var FindingKinds = []string{
	PoolExhausted,
	MissingTranscriptEmpty,
	MissingTranscriptKept,
	OrphanTrackerState,
	UnownedQueue,
	UnownedNotification,
	StaleCallsign,
	MigrationIncomplete,
}
