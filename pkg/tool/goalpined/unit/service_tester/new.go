package service_tester

import (
	"crypto/x509"
	"encoding/pem"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/unit/package_server_tester"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/face"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/option"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/service"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/unit/recording_indexer"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	root := filepath.Join(t.TempDir(), "packages")
	keys := t.TempDir()
	key := "test.rsa"
	system.WriteFile(
		filepath.Join(keys, key),
		pem.EncodeToMemory(
			&pem.Block{
				Type: "RSA PRIVATE KEY",
				Bytes: x509.MarshalPKCS1PrivateKey(
					package_server_tester.GenerateRSAKey(2048),
				),
			},
		),
		0o600,
	)
	i := recording_indexer.New()
	o := option.New()
	o.PackageRoot = root
	o.KeyDirectory = keys
	o.SignatureKey = key
	o.Indexer = func() face.Indexer { return i }

	return &Tester{
		Service:      service.New(o),
		Indexer:      i,
		PackageRoot:  root,
		KeyDirectory: keys,
		SignatureKey: key,
	}
}
