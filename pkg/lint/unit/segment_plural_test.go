package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/segment"
	"testing"
)

func TestSegmentsPluralInitialism(t *testing.T) {
	assert.Strings(t, []string{"down", "osds"}, segment.Segments("DownOSDs"))
}

func TestSegmentsBarePluralInitialism(t *testing.T) {
	assert.Strings(t, []string{"osds"}, segment.Segments("OSDs"))
}

func TestSegmentsInteriorPluralInitialism(t *testing.T) {
	assert.Strings(
		t,
		[]string{"num", "osds", "up"},
		segment.Segments("NumOSDsUp"),
	)
}

func TestSegmentsPluralIdentifiers(t *testing.T) {
	assert.Strings(t, []string{"user", "ids"}, segment.Segments("UserIDs"))
}

func TestSegmentsPluralLocators(t *testing.T) {
	assert.Strings(t, []string{"parse", "urls"}, segment.Segments("ParseURLs"))
}

func TestSegmentsInitialismBeforeWord(t *testing.T) {
	assert.Strings(t, []string{"osd", "map"}, segment.Segments("OSDMap"))
}

func TestSegmentsInitialismBeforeLowerWord(t *testing.T) {
	assert.Strings(t, []string{"os", "daemon"}, segment.Segments("OSDaemon"))
}

func TestSegmentsMixedCaseWord(t *testing.T) {
	assert.Strings(t, []string{"osd", "map"}, segment.Segments("OsdMap"))
}

func TestSegmentsProductNameKeepsSplit(t *testing.T) {
	assert.Strings(t, []string{"cent", "os"}, segment.Segments("CentOS"))
}
