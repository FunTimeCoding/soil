package unit

import (
	"crypto/x509"
	"encoding/pem"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/armor"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/authority_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/server_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/store_tester"
	"testing"
	"time"
)

func TestIntermediateInheritsRootUnderService(t *testing.T) {
	s, _ := store_tester.NewChain(t)
	defer s.Close()
	c := s.MustAuthority(constant.FixtureClusterAuthority).Material().Certificate
	assert.String(t, "Example Root CA", c.Issuer.CommonName)
}

func TestSecondRootConflicts(t *testing.T) {
	s, v := store_tester.NewChain(t)
	defer s.Close()
	_, e := v.CreateAuthority(server_tester.RootBody())
	assert.True(t, conflict.Is(e))
}

func TestIssuedLeafVerifiesAgainstTheChain(t *testing.T) {
	s, v := store_tester.NewChain(t)
	defer s.Close()
	r, key, e := v.IssueCertificate(
		server_tester.LeafBody(
			constant.FixtureCommonName,
			[]string{constant.FixtureHost},
		),
	)
	assert.Nil(t, e)
	assert.StringContains(t, "PRIVATE KEY", key)
	root, cluster := store_tester.ChainOf(s)
	assert.Nil(
		t,
		authority_tester.Verify(
			root,
			cluster,
			armor.DecodeCertificate([]byte(r.Certificate)),
		),
	)
}

func TestIssuedLeafKeyIsNotStored(t *testing.T) {
	s, v := store_tester.NewChain(t)
	defer s.Close()
	r, _, e := v.IssueCertificate(
		server_tester.LeafBody(
			constant.FixtureCommonName,
			[]string{constant.FixtureHost},
		),
	)
	assert.Nil(t, e)
	assert.String(t, "", s.MustBySerial(r.Serial).Key)
}

func TestForeignNameStillFailsThroughTheService(t *testing.T) {
	s, v := store_tester.NewChain(t)
	defer s.Close()
	r, _, e := v.IssueCertificate(
		server_tester.LeafBody(
			constant.FixtureImpostor,
			[]string{constant.FixtureForeignDomain},
		),
	)
	assert.Nil(t, e)
	root, cluster := store_tester.ChainOf(s)
	assert.Error(
		t,
		authority_tester.Verify(
			root,
			cluster,
			armor.DecodeCertificate([]byte(r.Certificate)),
		),
	)
}

func TestSignedRequestVerifiesAndStoresNoKey(t *testing.T) {
	s, v := store_tester.NewChain(t)
	defer s.Close()
	r, e := v.SignRequest(
		&server.SigningRequestBody{
			Authority: constant.FixtureClusterAuthority,
			Kind:      server.LeafKind(constant.KindServer),
			Request: authority_tester.NewSigningRequest(
				constant.FixtureRequestHost,
			),
		},
	)
	assert.Nil(t, e)
	assert.String(t, "", r.Key)
	root, cluster := store_tester.ChainOf(s)
	assert.Nil(
		t,
		authority_tester.Verify(
			root,
			cluster,
			armor.DecodeCertificate([]byte(r.Certificate)),
		),
	)
}

func TestRevocationListCarriesTheRevokedSerial(t *testing.T) {
	s, v := store_tester.NewChain(t)
	defer s.Close()
	r, _, e := v.IssueCertificate(
		server_tester.LeafBody(
			constant.FixtureCommonName,
			[]string{constant.FixtureHost},
		),
	)
	assert.Nil(t, e)
	assert.Nil(t, s.Revoke(r.Serial, time.Now()))
	b, f := v.RevocationList(constant.FixtureClusterAuthority)
	assert.Nil(t, f)
	block, _ := pem.Decode(b)
	list, g := x509.ParseRevocationList(block.Bytes)
	assert.Nil(t, g)
	assert.Integer(t, 1, len(list.RevokedCertificateEntries))
	assert.String(
		t,
		r.Serial,
		list.RevokedCertificateEntries[0].SerialNumber.Text(constant.SerialBase),
	)
}
