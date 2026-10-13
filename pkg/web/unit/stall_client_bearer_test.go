package unit

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"log"
	"os"
	"testing"
)

func TestStallClientBearerAuthReachesBothViews(t *testing.T) {
	s, seen := newHeaderServer(t)
	p := stall_client.New().BearerAuth("alfa").Build()
	defer p.Close()
	web.GetString(p.Client(), s.URL)
	assert.String(t, "Bearer alfa", *seen)
	*seen = ""
	web.GetString(p.Halting(), s.URL)
	assert.String(t, "Bearer alfa", *seen)
}

func TestStallClientVerboseNeverLogsBearer(t *testing.T) {
	s, seen := newHeaderServer(t)
	var b bytes.Buffer
	log.SetOutput(&b)
	defer log.SetOutput(os.Stderr)
	p := stall_client.New().Verbose().BearerAuth("alfa").Build()
	defer p.Close()
	web.GetString(p.Client(), s.URL)
	assert.String(t, "Bearer alfa", *seen)
	assert.StringContains(t, "REQUEST:", b.String())
	assert.StringNotContains(t, "alfa", b.String())
}
