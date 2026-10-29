package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/unit/mattermost_client_tester"
	"net/http"
	"sync"
	"testing"
)

func TestCloseAndRefreshDoNotRace(t *testing.T) {
	r := mattermost_client_tester.New(t, func(_ *http.ServeMux) {})
	var w sync.WaitGroup
	w.Go(
		func() {
			for range 20 {
				if e := r.Client.RefreshSocket(); e != nil {
					return
				}
			}
		},
	)
	r.Client.Close()
	w.Wait()
	assert.Error(t, r.Client.RefreshSocket())
	assert.True(t, r.Client.WebSocket() == nil)
}
