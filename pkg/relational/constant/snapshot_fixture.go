package constant

import "time"

const (
	FixtureLiteFile         = "writers.sqlite"
	FixtureOtherWriterStart = 300 * time.Millisecond
	FixtureOtherWriterWait  = 10 * time.Second
	FixtureEventCount       = "SELECT COUNT(*) FROM event"
	FixtureBeginDeferred    = "BEGIN DEFERRED"
	FixtureRollback         = "ROLLBACK"
)
