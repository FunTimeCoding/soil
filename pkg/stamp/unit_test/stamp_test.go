package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/stamp"
	"github.com/funtimecoding/soil/pkg/stamp/constant"
	"testing"
)

func TestStamp(t *testing.T) {
	s := stamp.New("a", "b", "c")
	assert.String(t, s.Version, "a")
	assert.String(t, s.GitHash, "b")
	assert.String(t, s.BuildDate, "c")
	d := stamp.New("", "", "")
	assert.String(t, d.Version, constant.DefaultVersion)
	assert.String(t, d.GitHash, constant.DefaultGitHash)
	assert.String(t, d.BuildDate, constant.DefaultDate)
}
