package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/authority"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/authority_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/store_tester"
	"testing"
	"time"
)

func TestAuthorityRoundTripKeepsSigningKey(t *testing.T) {
	s := store_tester.NewStore()
	defer s.Close()
	root := authority_tester.NewRoot()
	s.MustCreate(
		*record.New(constant.KindRoot, constant.RootAuthority, root.Material()),
	)
	restored := authority.New(
		s.MustAuthority(constant.RootAuthority).Material(),
	)
	cluster := authority_tester.NewCluster(restored)
	leaf := authority_tester.NewLeaf(
		cluster,
		constant.FixtureCommonName,
		[]string{constant.FixtureHost},
	)
	assert.Nil(t, authority_tester.Verify(restored, cluster, leaf.Certificate))
}

func TestMissingAuthorityIsNotAnError(t *testing.T) {
	s := store_tester.NewStore()
	defer s.Close()
	assert.Nil(t, s.MustAuthority(constant.RootAuthority))
}

func TestLeafKeyIsNotRetained(t *testing.T) {
	s := store_tester.NewStore()
	defer s.Close()
	cluster := authority_tester.NewCluster(authority_tester.NewRoot())
	leaf := authority_tester.NewLeaf(
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
		authority_tester.NewRoot().Material(),
	)
	assert.StringContains(t, "PRIVATE KEY", r.Key)
}

func TestRevokedAuthorityStopsResolving(t *testing.T) {
	s := store_tester.NewStore()
	defer s.Close()
	r := record.New(
		constant.KindRoot,
		constant.RootAuthority,
		authority_tester.NewRoot().Material(),
	)
	s.MustCreate(*r)
	errors.PanicOnError(s.Revoke(r.Serial, time.Now()))
	assert.Nil(t, s.MustAuthority(constant.RootAuthority))
}

func TestExpiringFindsOnlyPastTheHorizon(t *testing.T) {
	s := store_tester.NewStore()
	defer s.Close()
	cluster := authority_tester.NewCluster(authority_tester.NewRoot())
	leaf := authority_tester.NewLeaf(
		cluster,
		constant.FixtureCommonName,
		[]string{constant.FixtureHost},
	)
	s.MustCreate(*record.New(constant.KindServer, "", leaf))
	near, e := s.Expiring(time.Now().AddDate(0, 1, 0))
	errors.PanicOnError(e)
	assert.Integer(t, 0, len(near))
	far, f := s.Expiring(time.Now().AddDate(2, 0, 0))
	errors.PanicOnError(f)
	assert.Integer(t, 1, len(far))
}
