package constant

import "time"

const (
	FixtureQuietRoute   = "/quiet"
	FixtureBusyRoute    = "/busy"
	FixtureStalledRoute = "/stalled"
	FixtureHangRoute    = "/hang"
	FixtureHeavyRoute   = "/heavy"
	FixtureUnloadRoute  = "/unload"
	FixturePopupRoute   = "/popup"
	FixturePingRoute    = "/ping"

	FixtureAbsentTab = "nowhere"

	FixtureHeavyRowCount = 4000
)

const (
	FixtureWatchBuffer      = 16
	FixtureSharedCycleCount = 10

	FixtureCloseSettlePeriod = 1500 * time.Millisecond
	FixtureEventTimeout      = 5 * time.Second
)
