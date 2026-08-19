package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
)

func Certificate(r *record.Record) *server.CertificateResponse {
	result := &server.CertificateResponse{
		Serial:      r.Serial,
		Kind:        server.CertificateKind(r.Kind),
		CommonName:  r.CommonName,
		Issuer:      r.Issuer,
		Certificate: r.Certificate,
		NotBefore:   r.Start,
		NotAfter:    r.End,
		RevokedAt:   r.Revoked,
	}

	if h := r.Material().Certificate.DNSNames; len(h) > 0 {
		result.Host = &h
	}

	return result
}
