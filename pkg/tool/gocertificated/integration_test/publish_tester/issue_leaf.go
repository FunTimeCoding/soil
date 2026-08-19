package publish_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	"testing"
)

func (o *Tester) IssueLeaf(t *testing.T) {
	t.Helper()
	host := []string{constant.FixtureHost}
	_, _, e := o.Server.Service.IssueCertificate(
		&server.CertificateBody{
			Authority:  constant.FixtureClusterAuthority,
			Kind:       server.LeafKind(constant.KindServer),
			CommonName: constant.FixtureCommonName,
			Host:       &host,
		},
	)
	assert.FatalOnError(t, e)
}
