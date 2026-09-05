package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/armor"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/authority"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/keystore"
	"os"
	"path/filepath"
	"testing"
)

func TestCertificateSurvivesArmorRoundTrip(t *testing.T) {
	c := newRoot().Material().Certificate
	assert.String(
		t,
		c.Subject.CommonName,
		armor.DecodeCertificate(armor.MarshalCertificate(c)).Subject.CommonName,
	)
}

func TestKeySignsAfterArmorRoundTrip(t *testing.T) {
	root := newRoot()
	directory := filepath.Join(t.TempDir(), "material")
	keystore.Write(directory, root.Material())
	restored := authority.New(keystore.Read(directory))
	cluster := newCluster(restored)
	leaf := newLeaf(
		cluster,
		constant.FixtureCommonName,
		[]string{constant.FixtureHost},
	)
	assert.Nil(t, verify(restored, cluster, leaf.Certificate))
}

func TestKeystoreRoundTripPreservesSubject(t *testing.T) {
	directory := filepath.Join(t.TempDir(), "material")
	keystore.Write(directory, newRoot().Material())
	assert.String(
		t,
		"Example Root CA",
		keystore.Read(directory).Certificate.Subject.CommonName,
	)
}

func TestKeyFileIsNotReadableByOthers(t *testing.T) {
	directory := filepath.Join(t.TempDir(), "material")
	keystore.Write(directory, newRoot().Material())
	i, e := os.Stat(filepath.Join(directory, constant.KeyFile))

	if e != nil {
		panic(e)
	}

	assert.String(t, "-rw-------", i.Mode().String())
}
