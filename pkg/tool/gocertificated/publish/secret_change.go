package publish

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gosecret"
)

func (p *Publisher) secretChange(r *record.Record) ([]*Change, error) {
	payload := map[string]string{
		constant.SecretCertificateKey: r.Certificate,
		constant.SecretKeyKey:         r.Key,
	}
	manifest, e := gosecret.ReplacePayload(
		[]byte(secretSkeleton(p.secretPath)),
		payload,
	)

	if e != nil {
		return nil, e
	}

	reason := join.Space(r.Kind, r.CommonName)

	return []*Change{
		NewChange(p.secretPath, reason, string(manifest)),
		NewChange(
			gosecret.GetDecodedPath(p.secretPath),
			reason,
			gosecret.DecodedContent(payload),
		),
	}, nil
}
