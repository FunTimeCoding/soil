package server_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
)

func LeafBody(
	common string,
	host []string,
) *server.CertificateBody {
	return &server.CertificateBody{
		Authority:  constant.FixtureClusterAuthority,
		Kind:       server.LeafKind(constant.KindServer),
		CommonName: common,
		Host:       &host,
	}
}
