package constant

import "errors"

const FrontMatterDelimiter = "---"
const (
	RegionStartFormat = "<!-- %s:start -->"
	RegionEndFormat   = "<!-- %s:end -->"

	RegionRepeated  = "region %q: %s appears %d times"
	RegionUnpaired  = "region %q: %s has no partner"
	RegionReversed  = "region %q: end marker comes before start marker"
	RegionStartWord = "start marker"
	RegionEndWord   = "end marker"

	FixtureRegion = "alfa"
)

var ErrorRegionMissing = errors.New("region markers missing")
