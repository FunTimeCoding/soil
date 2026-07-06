package notation

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"testing"
)

func TestDecode(t *testing.T) {
	var actual []int
	assert.FatalOnError(t, Decode("[1]", &actual))
	assert.Any(t, []int{1}, actual)
}

func TestDecodeStrict(t *testing.T) {
	var actual []int
	MustDecode("[1]", &actual, false)
	assert.Any(t, []int{1}, actual)
}

type User struct {
	Name    string         `json:"name"`
	Unknown map[string]any `json:"-"`
}

func (u *User) UnmarshalJSON(b []byte) error {
	// Prevent recursion
	type Alias User
	v := (*Alias)(u)

	return UnmarshalUnknown(b, v, UnknownField)
}

func (u *User) UnknownField() map[string]any {
	return u.Unknown
}

func TestUnknown(t *testing.T) {
	raw := `{
		"name": "jdoe",
		"department": "Development",
		"location": "Earth",
		"skills": ["Go", "Kubernetes"]
	}`
	var u User
	errors.PanicOnError(json.Unmarshal([]byte(raw), &u))
	assert.String(t, "jdoe", u.Name)
	assert.Any(t, "Development", u.Unknown["department"])
	assert.Any(t, "Earth", u.Unknown["location"])
	assert.Any(t, []any{"Go", "Kubernetes"}, u.Unknown["skills"])
}
