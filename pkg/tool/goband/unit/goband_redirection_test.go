package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/base"
	"testing"
)

func TestRedirectionDryRunNamesThePendingWrites(t *testing.T) {
	s := base.New(t)
	s.Alfa.SetRedirection(false, 3)
	output, _, code := run(t, s, "redirection", "--instance", "alfa")
	assert.Integer(t, 0, code)
	assert.StringContains(
		t,
		"Would enable listener, screen service on alfa. Pass --yes to execute.",
		output,
	)
	assert.Count(t, 0, s.Alfa.Writes())
}

func TestRedirectionWithYesEnablesWhatIsMissing(t *testing.T) {
	s := base.New(t)
	s.Alfa.SetRedirection(false, 3)
	output, _, code := run(t, s, "redirection", "--instance", "alfa", "--yes")
	assert.Integer(t, 0, code)
	assert.StringContains(t, "Enabled listener, screen service on alfa", output)
	assert.Strings(t, []string{"listener", "screen service"}, s.Alfa.Writes())
}

func TestRedirectionSaysWhenNothingIsMissing(t *testing.T) {
	s := base.New(t)
	s.Alfa.SetRedirection(true, constant.ScreenServiceOffline)
	output, _, code := run(t, s, "redirection", "--instance", "alfa", "--yes")
	assert.Integer(t, 0, code)
	assert.StringContains(t, "Redirection already enabled on alfa", output)
	assert.Count(t, 0, s.Alfa.Writes())
}
