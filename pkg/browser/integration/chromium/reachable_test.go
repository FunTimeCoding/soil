//go:build browser

package chromium

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/browser/integration/chromium/base"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/browser"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/face"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/mock_client"
	"testing"
)

func TestBrowserAndMockHoldChromiumSource(t *testing.T) {
	s := base.New(t)

	for _, f := range []face.ChromiumSource{
		browser.New(s.Client),
		mock_client.New(t),
	} {
		assert.NotNil(t, f)
	}
}

func TestReachableAnswersWhileBrowserRuns(t *testing.T) {
	s := base.New(t)
	assert.FatalOnError(t, s.Client.Reachable(context.Background()))
}
