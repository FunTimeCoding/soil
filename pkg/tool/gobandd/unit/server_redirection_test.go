package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	band "github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/base"
	"testing"
)

func TestEnableRedirectionWritesOnlyWhatIsMissing(t *testing.T) {
	s := base.New(t)
	s.Alfa.SetRedirection(false, 3)
	body := client.EnableRedirectionJSONRequestBody{Instance: new("alfa")}
	r, e := s.Client.EnableRedirectionWithResponse(t.Context(), body)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, r.StatusCode())
	assert.Strings(t, []string{"listener", "screen service"}, r.JSON200.Enabled)
	again, f := s.Client.EnableRedirectionWithResponse(t.Context(), body)
	assert.FatalOnError(t, f)
	assert.Count(t, 0, again.JSON200.Enabled)
	assert.Strings(t, []string{"listener", "screen service"}, s.Alfa.Writes())
}

func TestEnableRedirectionLeavesAnEnabledButOfflineScreenServiceAlone(
	t *testing.T,
) {
	s := base.New(t)
	s.Alfa.SetRedirection(true, band.ScreenServiceOffline)
	r, e := s.Client.EnableRedirectionWithResponse(
		t.Context(),
		client.EnableRedirectionJSONRequestBody{Instance: new("alfa")},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, r.JSON200.Enabled)
	assert.Count(t, 0, s.Alfa.Writes())
}
