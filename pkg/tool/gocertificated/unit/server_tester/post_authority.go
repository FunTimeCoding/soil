package server_tester

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
)

func (o *Tester) PostAuthority(
	c *client.ClientWithResponses,
	b *server.AuthorityBody,
) *client.PostAuthorityResponse {
	o.t.Helper()
	result, e := c.PostAuthorityWithBodyWithResponse(
		o.t.Context(),
		webConstant.Object,
		bytes.NewReader(notation.Marshal(b)),
	)
	assert.FatalOnError(o.t, e)

	return result
}
