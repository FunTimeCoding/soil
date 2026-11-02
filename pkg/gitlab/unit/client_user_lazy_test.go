package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"sync/atomic"
	"testing"
)

func TestTheUserIsFetchedOnFirstNeed(t *testing.T) {
	var users atomic.Int64
	var assignee atomic.Value
	r := client_tester.New(
		t,
		func(m *http.ServeMux) {
			m.HandleFunc(
				"/api/v4/user",
				func(
					w http.ResponseWriter,
					_ *http.Request,
				) {
					users.Add(1)
					web.ObjectHeader(w)
					writer.Print(w, `{"id": 42, "username": "tester"}`)
				},
			)
			m.HandleFunc(
				"/api/v4/merge_requests",
				func(
					w http.ResponseWriter,
					q *http.Request,
				) {
					assignee.Store(q.URL.Query().Get("assignee_id"))
					web.ObjectHeader(w)
					writer.Print(w, `[]`)
				},
			)
		},
	)
	assert.Integer(t, 0, int(users.Load()))
	_, e := r.Client.AssignedMergeRequests(true)
	assert.FatalOnError(t, e)
	_, e = r.Client.AssignedMergeRequests(true)
	assert.FatalOnError(t, e)
	assert.Integer(t, 1, int(users.Load()))
	assert.Any(t, "42", assignee.Load())
}

func TestAUserLookupRefusedIsAnError(t *testing.T) {
	r := client_tester.New(
		t,
		func(m *http.ServeMux) {
			m.HandleFunc(
				"/api/v4/user",
				func(
					w http.ResponseWriter,
					_ *http.Request,
				) {
					web.ObjectHeader(w)
					w.WriteHeader(http.StatusUnauthorized)
					writer.Print(w, `{"message": "401 Unauthorized"}`)
				},
			)
			m.HandleFunc(
				"/api/v4/merge_requests",
				func(
					w http.ResponseWriter,
					_ *http.Request,
				) {
					web.ObjectHeader(w)
					writer.Print(w, `[]`)
				},
			)
		},
	)
	_, e := r.Client.ReviewingMergeRequests(true)
	assert.Error(t, e)
}
