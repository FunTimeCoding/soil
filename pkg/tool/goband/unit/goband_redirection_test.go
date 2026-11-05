package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	bandConstant "github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goband/plan"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/base"
	"testing"
)

func TestRedirectionDryRunNamesThePendingWrites(t *testing.T) {
	s := base.New(t)
	s.Alfa.SetRedirection(false, 3)
	output, _, code := run(t, s, "redirection", "--instance", "alfa")
	assert.Integer(t, 0, code)
	var p plan.Redirection
	notation.MustDecode(output, &p, false)
	assert.String(t, "alfa", p.Instance)
	assert.Strings(t, []string{"listener", "screen service"}, p.Pending)
	assert.Count(t, 0, s.Alfa.Writes())
}

func TestRedirectionWithYesEnablesWhatIsMissing(t *testing.T) {
	s := base.New(t)
	s.Alfa.SetRedirection(false, 3)
	output, _, code := run(t, s, "redirection", "--instance", "alfa", "--yes")
	assert.Integer(t, 0, code)
	var r client.RedirectionResult
	notation.MustDecode(output, &r, false)
	assert.String(t, "alfa", r.Instance)
	assert.Strings(t, []string{"listener", "screen service"}, r.Enabled)
	assert.Strings(t, []string{"listener", "screen service"}, s.Alfa.Writes())
}

func TestRedirectionSaysWhenNothingIsMissing(t *testing.T) {
	s := base.New(t)
	s.Alfa.SetRedirection(true, bandConstant.ScreenServiceOffline)
	output, _, code := run(t, s, "redirection", "--instance", "alfa", "--yes")
	assert.Integer(t, 0, code)
	var p plan.Redirection
	notation.MustDecode(output, &p, false)
	assert.String(t, "alfa", p.Instance)
	assert.Count(t, 0, p.Pending)
	assert.Count(t, 0, s.Alfa.Writes())
}
