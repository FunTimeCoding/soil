package unit

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func postAuthority(
	t *testing.T,
	c *client.ClientWithResponses,
	b *server.AuthorityBody,
) *client.PostAuthorityResponse {
	t.Helper()
	result, e := c.PostAuthorityWithBodyWithResponse(
		t.Context(),
		webConstant.Object,
		bytes.NewReader(notation.Marshal(b)),
	)
	assert.FatalOnError(t, e)

	return result
}
