package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/authority"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"testing"
	"time"
)

func newStore() *store.Store {
	return store.New(lite.NewMemory())
}

func TestAuthorityRoundTripKeepsSigningKey(t *testing.T) {
	s := newStore()
	defer s.Close()
	root := newRoot()
	s.MustCreate(
		*record.New(constant.KindRoot, constant.RootAuthority, root.Material()),
	)
	restored := authority.New(
		s.MustAuthority(constant.RootAuthority).Material(),
	)
	cluster := newCluster(restored)
	leaf := newLeaf(
		cluster,
		constant.FixtureCommonName,
		[]string{constant.FixtureHost},
	)
	assert.Nil(t, verify(restored, cluster, leaf.Certificate))
}

func TestMissingAuthorityIsNotAnError(t *testing.T) {
	s := newStore()
	defer s.Close()
	assert.Nil(t, s.MustAuthority(constant.RootAuthority))
}

func TestLeafKeyIsNotRetained(t *testing.T) {
	s := newStore()
	defer s.Close()
	cluster := newCluster(newRoot())
	leaf := newLeaf(
		cluster,
		constant.FixtureCommonName,
		[]string{constant.FixtureHost},
	)
	r := record.New(constant.KindServer, "", leaf)
	assert.String(t, "", r.Key)
}

func TestAuthorityKeyIsRetained(t *testing.T) {
	r := record.New(
		constant.KindRoot,
		constant.RootAuthority,
		newRoot().Material(),
	)
	assert.StringContains(t, "PRIVATE KEY", r.Key)
}

func TestRevokedAuthorityStopsResolving(t *testing.T) {
	s := newStore()
	defer s.Close()
	r := record.New(
		constant.KindRoot,
		constant.RootAuthority,
		newRoot().Material(),
	)
	s.MustCreate(*r)
	errorless(s.Revoke(r.Serial, time.Now()))
	assert.Nil(t, s.MustAuthority(constant.RootAuthority))
}

func TestExpiringFindsOnlyPastTheHorizon(t *testing.T) {
	s := newStore()
	defer s.Close()
	cluster := newCluster(newRoot())
	leaf := newLeaf(
		cluster,
		constant.FixtureCommonName,
		[]string{constant.FixtureHost},
	)
	s.MustCreate(*record.New(constant.KindServer, "", leaf))
	near, e := s.Expiring(time.Now().AddDate(0, 1, 0))
	errorless(e)
	assert.Integer(t, 0, len(near))
	far, f := s.Expiring(time.Now().AddDate(2, 0, 0))
	errorless(f)
	assert.Integer(t, 1, len(far))
}

func errorless(e error) {
	if e != nil {
		panic(e)
	}
}
