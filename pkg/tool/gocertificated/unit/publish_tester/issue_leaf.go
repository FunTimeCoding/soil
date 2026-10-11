package publish_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
)

func (o *Tester) IssueLeaf() {
	o.t.Helper()
	host := []string{constant.FixtureHost}
	_, _, e := o.Server.Service.IssueCertificate(
		&server.CertificateBody{
			Authority:  constant.FixtureClusterAuthority,
			Kind:       server.LeafKind(constant.KindServer),
			CommonName: constant.FixtureCommonName,
			Host:       &host,
		},
	)
	assert.FatalOnError(o.t, e)
}
