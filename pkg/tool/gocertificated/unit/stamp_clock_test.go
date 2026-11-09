package unit

import (
	"crypto/x509"
	"encoding/pem"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/filter"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/authority_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/server_tester"
	"testing"
)

func TestRootValidityRunsOnServiceClock(t *testing.T) {
	_, root := clockedChain(t)
	assert.True(t, root.Start.Equal(windowNow().Add(-constant.ClockSkew)))
}

func TestIntermediateValidityRunsOnServiceClock(t *testing.T) {
	v, _ := clockedChain(t)
	f := filter.New()
	f.Kind = string(constant.KindIntermediate)
	intermediate, e := v.Certificates(f)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, intermediate)
	assert.True(
		t,
		intermediate[0].Start.Equal(windowNow().Add(-constant.ClockSkew)),
	)
}

func TestIssuedValidityRunsOnServiceClock(t *testing.T) {
	v, _ := clockedChain(t)
	r, _, e := v.IssueCertificate(
		server_tester.LeafBody(
			constant.FixtureCommonName,
			[]string{constant.FixtureHost},
		),
	)
	assert.FatalOnError(t, e)
	assert.True(t, r.Start.Equal(windowNow().Add(-constant.ClockSkew)))
}

func TestSignedValidityRunsOnServiceClock(t *testing.T) {
	v, _ := clockedChain(t)
	r, e := v.SignRequest(
		&server.SigningRequestBody{
			Authority: constant.FixtureClusterAuthority,
			Kind:      server.LeafKind(constant.KindServer),
			Request: authority_tester.NewSigningRequest(
				constant.FixtureRequestHost,
			),
		},
	)
	assert.FatalOnError(t, e)
	assert.True(t, r.Start.Equal(windowNow().Add(-constant.ClockSkew)))
}

func TestRevokeStampRunsOnServiceClock(t *testing.T) {
	v, _ := clockedChain(t)
	r, _, e := v.IssueCertificate(
		server_tester.LeafBody(
			constant.FixtureCommonName,
			[]string{constant.FixtureHost},
		),
	)
	assert.FatalOnError(t, e)
	revoked, f := v.Revoke(r.Serial)
	assert.FatalOnError(t, f)
	assert.FatalNotNil(t, revoked.Revoked)
	assert.True(t, revoked.Revoked.Equal(windowNow()))
}

func TestRevocationListRunsOnServiceClock(t *testing.T) {
	v, _ := clockedChain(t)
	b, e := v.RevocationList(constant.FixtureClusterAuthority)
	assert.FatalOnError(t, e)
	block, _ := pem.Decode(b)
	list, f := x509.ParseRevocationList(block.Bytes)
	assert.FatalOnError(t, f)
	assert.True(t, list.ThisUpdate.Equal(windowNow()))
}

func TestPublishStampRunsOnServiceClock(t *testing.T) {
	v, _ := clockedChain(t)
	_, _, e := v.Publish(t.Context())
	assert.FatalOnError(t, e)
	published, f := v.Certificates(filter.New())
	assert.FatalOnError(t, f)
	assert.Count(t, 2, published)

	for _, r := range published {
		assert.FatalNotNil(t, r.Published)
		assert.True(t, r.Published.Equal(windowNow()))
	}
}
