package authority

import (
	"crypto/rand"
	"crypto/x509"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/armor"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"time"
)

func (a *Authority) RevocationList(v []x509.RevocationListEntry) []byte {
	now := time.Now()
	t := &x509.RevocationList{}
	t.Number = newSerial()
	t.ThisUpdate = now
	t.NextUpdate = now.AddDate(0, 0, constant.RevocationListDay)
	t.RevokedCertificateEntries = v
	b, e := x509.CreateRevocationList(
		rand.Reader,
		t,
		a.material.Certificate,
		a.material.Key,
	)
	errors.PanicOnError(e)

	return armor.MarshalRevocationList(b)
}
