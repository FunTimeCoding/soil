package unit

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/notation/fixture"
	"testing"
)

func TestDecode(t *testing.T) {
	var actual []int
	assert.FatalOnError(t, notation.Decode("[1]", &actual))
	assert.Any(t, []int{1}, actual)
}

func TestDecodeStrict(t *testing.T) {
	var actual []int
	notation.MustDecode("[1]", &actual, false)
	assert.Any(t, []int{1}, actual)
}

func TestUnknown(t *testing.T) {
	raw := `{
		"name": "jdoe",
		"department": "Development",
		"location": "Earth",
		"skills": ["Go", "Kubernetes"]
	}`
	var u fixture.User
	errors.PanicOnError(json.Unmarshal([]byte(raw), &u))
	assert.String(t, "jdoe", u.Name)
	assert.Any(t, "Development", u.Unknown["department"])
	assert.Any(t, "Earth", u.Unknown["location"])
	assert.Any(t, []any{"Go", "Kubernetes"}, u.Unknown["skills"])
}
