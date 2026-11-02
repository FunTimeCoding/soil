package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/telemetry"
	"github.com/funtimecoding/soil/pkg/telemetry/constant"
	"github.com/funtimecoding/soil/pkg/telemetry/record"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestARefusedTelemetryDaemonCostsTheCommandNothing(t *testing.T) {
	s := httptest.NewServer(http.NotFoundHandler())
	host, port := upstream_tester.HostPort(t, s.URL)
	s.Close()
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

func TestAStalledTelemetryDaemonHoldsTheCommandOnlyToItsTimeout(t *testing.T) {
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
	host, port := upstream_tester.HostPort(t, s.URL)
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
