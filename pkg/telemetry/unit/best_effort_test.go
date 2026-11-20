package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/telemetry"
	"github.com/funtimecoding/soil/pkg/telemetry/constant"
	"github.com/funtimecoding/soil/pkg/telemetry/record"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRefusedTelemetryDaemonCostsCommandNothing(t *testing.T) {
	host, port := upstream_server.HostPort(
		t,
		outage_tester.New(t).RefusedAddress(),
	)
	c := telemetry.New(host, port, true).WithTimeout(
		constant.CommandLineTimeout,
	)
	defer c.Close()
	started := time.Now()
	c.Record(
		record.NewDomain(
			"alfa_list",
			constant.CommandLine,
			constant.User,
			constant.Success,
		),
	)
	c.Flush()
	assert.True(t, time.Since(started) < constant.CommandLineTimeout)
}

func TestStalledTelemetryDaemonHoldsCommandOnlyToItsTimeout(t *testing.T) {
	release := make(chan struct{})
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				http.ResponseWriter,
				*http.Request,
			) {
				<-release
			},
		),
	)
	defer s.Close()
	defer close(release)
	host, port := upstream_server.HostPort(t, s.URL)
	c := telemetry.New(host, port, true).WithTimeout(
		constant.FixtureStallTimeout,
	)
	defer c.Close()
	c.Record(
		record.NewDomain(
			"alfa_list",
			constant.CommandLine,
			constant.User,
			constant.Success,
		),
	)
	flushed := make(chan struct{})
	go func() {
		c.Flush()
		close(flushed)
	}()

	select {
	case <-flushed:
	case <-time.After(constant.FixtureStallBound):
		t.Fatal("telemetry held the command past its timeout")
	}
}
