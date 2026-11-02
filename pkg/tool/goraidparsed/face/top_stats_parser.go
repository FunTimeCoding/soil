package face

import "time"

type TopStatsParser interface {
	Parse(
		workdir string,
		date time.Time,
	) error
	Import(
		template string,
		workdir string,
	) error
}
