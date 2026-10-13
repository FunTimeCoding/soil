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

func TestStallClientBasicAuthReachesBothViews(t *testing.T) {
	s, seen := newBasicAuthServer(t)
	p := stall_client.New().BasicAuth("alfa", "bravo").Build()
	defer p.Close()
	web.GetString(p.Client(), s.URL)
	assert.String(t, "alfa:bravo", *seen)
	*seen = ""
	web.GetString(p.Halting(), s.URL)
	assert.String(t, "alfa:bravo", *seen)
}

func TestStallClientVerboseNeverLogsCredentials(t *testing.T) {
	s, seen := newBasicAuthServer(t)
	var b bytes.Buffer
	log.SetOutput(&b)
	defer log.SetOutput(os.Stderr)
	p := stall_client.New().Verbose().BasicAuth("alfa", "bravo").Build()
	defer p.Close()
	web.GetString(p.Client(), s.URL)
	assert.String(t, "alfa:bravo", *seen)
	assert.StringContains(t, "REQUEST:", b.String())
	assert.StringNotContains(t, "Authorization", b.String())
}
