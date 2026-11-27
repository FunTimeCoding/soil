package constant

import "time"

const (
	ChromiumFixtureQuietRoute   = "/quiet"
	ChromiumFixtureBusyRoute    = "/busy"
	ChromiumFixtureStalledRoute = "/stalled"
	ChromiumFixtureHangRoute    = "/hang"
	ChromiumFixtureHeavyRoute   = "/heavy"
	ChromiumFixtureUnloadRoute  = "/unload"
	ChromiumFixturePopupRoute   = "/popup"
	ChromiumFixturePingRoute    = "/ping"

	ChromiumFixtureAbsentTab = "nowhere"
	ChromiumFixtureStartPage = "about:blank"

	ChromiumFixtureHeavyRowCount    = 4000
	ChromiumFixtureWatchBuffer      = 16
	ChromiumFixtureSharedCycleCount = 10

	ChromiumFixtureCloseSettlePeriod = 1500 * time.Millisecond
	ChromiumFixtureEventTimeout      = 5 * time.Second
	ChromiumFixtureCallTimeout       = 500 * time.Millisecond
	ChromiumFixtureHungPromise       = "new Promise(() => {})"
	ChromiumFixtureSum               = "1 + 1"
	ChromiumFixtureTimer             = "new Promise(r => setTimeout(() => r('timer'), 10))"
)
