package server_tester

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/client"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
)

func (o *Tester) PostCertificate(
	c *client.ClientWithResponses,
) *client.PostCertificateResponse {
	o.t.Helper()
	result, e := c.PostCertificateWithBodyWithResponse(
		o.t.Context(),
		webConstant.Object,
		bytes.NewReader(
			notation.Marshal(
				LeafBody(
					constant.FixtureCommonName,
					[]string{constant.FixtureHost},
				),
			),
		),
	)
	assert.FatalOnError(o.t, e)

	return result
}
