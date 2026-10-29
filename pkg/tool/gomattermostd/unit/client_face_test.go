package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/client_tester"
	"net/http"
	"testing"
)

func TestTheClientAndItsMockShareTheFace(t *testing.T) {
	m := mock_client.New()
	t.Cleanup(m.Close)
	faces := []face.Client{
		client_tester.New(t, func(_ *http.ServeMux) {}).Client,
		m,
	}
	assert.Count(t, 2, faces)
}
