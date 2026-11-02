package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/firefox"
	"github.com/funtimecoding/soil/pkg/firefox/face"
	"github.com/funtimecoding/soil/pkg/firefox/unit/fake_extension"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/option"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/unit/base"
	"go.uber.org/goleak"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTheDaemonClosesItsClientOnce(t *testing.T) {
	s := base.New(t)
	s.Daemon.Close()
	s.Daemon.Close()
	assert.Integer(t, 1, s.MockClient.Closed)
}

func TestTheDaemonLeavesNothingRunningWithTheExtensionConnected(t *testing.T) {
	before := goleak.IgnoreCurrent()
	c := firefox.New()
	o := option.New()
	o.DownloadDirectory = t.TempDir()
	o.Client = func() face.FirefoxClient { return c }
	d := daemon.New(o)
	m := http.NewServeMux()
	d.MountBridge(m)
	s := httptest.NewServer(m)
	x := fake_extension.Connect(t, s.URL)
	defer x.Disconnect()
	deadline := time.Now().Add(5 * time.Second)

	for _, e := c.Tabs(); unavailable.Is(e); _, e = c.Tabs() {
		if time.Now().After(deadline) {
			t.Fatal("extension never connected")
		}

		time.Sleep(10 * time.Millisecond)
	}

	d.Close()
	s.Close()
	goleak.VerifyNone(t, before)
}
