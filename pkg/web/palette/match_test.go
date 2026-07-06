package palette

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestMatchExact(t *testing.T) {
	score, positions := Match("characters", "Characters")
	assert.True(t, score > 0)
	assert.Integer(t, 10, len(positions))
}

func TestMatchPrefix(t *testing.T) {
	score, positions := Match("char", "Characters")
	assert.True(t, score > 0)
	assert.Integer(t, 4, len(positions))
	assert.Integer(t, 0, positions[0])
	assert.Integer(t, 1, positions[1])
	assert.Integer(t, 2, positions[2])
	assert.Integer(t, 3, positions[3])
}

func TestMatchWordBoundaries(t *testing.T) {
	score, positions := Match("cp", "Create project")
	assert.True(t, score > 0)
	assert.Integer(t, 0, positions[0])
	assert.Integer(t, 7, positions[1])
}

func TestMatchCaseInsensitive(t *testing.T) {
	score, _ := Match("CREATE", "Create project")
	assert.True(t, score > 0)
}

func TestMatchNoMatch(t *testing.T) {
	score, _ := Match("xyz", "Create project")
	assert.Integer(t, -1, score)
}

func TestMatchEmptyPattern(t *testing.T) {
	score, _ := Match("", "Characters")
	assert.Integer(t, 0, score)
}

func TestMatchPatternLongerThanText(t *testing.T) {
	score, _ := Match("characters list", "char")
	assert.Integer(t, -1, score)
}

func TestMatchConsecutiveBeatsScattered(t *testing.T) {
	consecutive, _ := Match("dash", "Dashboard")
	scattered, _ := Match("dash", "Deploy at shared host")
	assert.True(t, consecutive > scattered)
}

func TestMatchWordBoundaryBeatsMiddle(t *testing.T) {
	boundary, _ := Match("sb", "Start build")
	middle, _ := Match("sb", "Testbed")
	assert.True(t, boundary > middle)
}

func TestMatchPrefixBeatsLater(t *testing.T) {
	prefix, _ := Match("cr", "Create project")
	later, _ := Match("cr", "search results")
	assert.True(t, prefix > later)
}

func TestMatchAcronymStyle(t *testing.T) {
	score, positions := Match("pd", "Push deploy")
	assert.True(t, score > 0)
	assert.Integer(t, 0, positions[0])
	assert.Integer(t, 5, positions[1])
}

func TestMatchNonContiguous(t *testing.T) {
	score, _ := Match("cpr", "Create project")
	assert.True(t, score > 0)
}

func TestMatchPositionsAreValid(t *testing.T) {
	_, positions := Match("sl", "Search logs")
	assert.Integer(t, 2, len(positions))

	for i := 1; i < len(positions); i++ {
		assert.True(t, positions[i] > positions[i-1])
	}
}
