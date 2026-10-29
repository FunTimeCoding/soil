package unit

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/server_tester"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func postCertificate(
	t *testing.T,
	c *client.ClientWithResponses,
) *client.PostCertificateResponse {
	t.Helper()
	result, e := c.PostCertificateWithBodyWithResponse(
		t.Context(),
		webConstant.Object,
		bytes.NewReader(
			notation.Marshal(
				server_tester.LeafBody(
					constant.FixtureCommonName,
					[]string{constant.FixtureHost},
				),
			),
		),
	)
	assert.FatalOnError(t, e)

	return result
}
