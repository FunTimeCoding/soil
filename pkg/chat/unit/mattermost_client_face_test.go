package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/face"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/mock_client"
	"github.com/funtimecoding/soil/pkg/chat/unit/mattermost_client_tester"
	"net/http"
	"testing"
)

func TestTheClientAndItsMockShareTheFace(t *testing.T) {
	r := mattermost_client_tester.New(t, func(_ *http.ServeMux) {})
	faces := []face.Client{r.Client, mock_client.New("me")}
	assert.Count(t, 2, faces)
}
