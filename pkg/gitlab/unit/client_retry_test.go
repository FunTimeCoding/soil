package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"sync/atomic"
	"testing"
)

func TestAWriteMeetingABadGatewayIsSentOnce(t *testing.T) {
	var sent atomic.Int64
	r := client_tester.New(
		t,
		func(m *http.ServeMux) {
			m.HandleFunc(
				"/api/v4/projects/7/pipeline",
				func(
					w http.ResponseWriter,
					_ *http.Request,
				) {
					sent.Add(1)
					web.ObjectHeader(w)
					w.WriteHeader(http.StatusBadGateway)
				},
			)
		},
		gitlab.WithRetryMax(1),
	)
	_, e := r.Client.CreatePipeline(7, "main", nil)
	assert.Error(t, e)
	assert.Integer(t, 1, int(sent.Load()))
}

func TestAReadMeetingABadGatewayIsRetried(t *testing.T) {
	var sent atomic.Int64
	r := client_tester.New(
		t,
		func(m *http.ServeMux) {
			m.HandleFunc(
				"/api/v4/projects/7",
				func(
					w http.ResponseWriter,
					_ *http.Request,
				) {
					sent.Add(1)
					web.ObjectHeader(w)
					w.WriteHeader(http.StatusBadGateway)
				},
			)
		},
		gitlab.WithRetryMax(1),
	)
	_, e := r.Client.Project(7)
	assert.Error(t, e)
	assert.Integer(t, 2, int(sent.Load()))
}
