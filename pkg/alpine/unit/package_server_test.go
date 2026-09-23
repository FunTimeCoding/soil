package unit

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/alpine/package_server"
	"github.com/funtimecoding/soil/pkg/alpine/unit/package_server_tester"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"maps"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

func TestSign(t *testing.T) {
	d := filepath.Join("../../..", "tmp", "alpine-signing-test")
	system.MakeDirectory(d)
	unsignedPath := filepath.Join(d, "test-unsigned.apk")
	package_server_tester.CreateTestPackage(d, unsignedPath)
	privateKey := package_server_tester.GenerateRSAKey(2048)
	unsignedSegments := package_server_tester.SplitArchive(unsignedPath)

	if len(unsignedSegments) != 2 {
		t.Fatalf(
			"unsigned package should have 2 segments, got %d",
			len(unsignedSegments),
		)
	}

	package_server.SignPackageWithKey(unsignedPath, privateKey, "test.rsa")
	signedSegments := package_server_tester.SplitArchive(unsignedPath)

	if len(signedSegments) != 3 {
		t.Fatalf(
			"signed package should have 3 segments, got %d",
			len(signedSegments),
		)
	}

	signatureSegment := signedSegments[0]
	signatureFiles := package_server_tester.ReadTarGz(signatureSegment)
	var signature []byte

	for name, content := range signatureFiles {
		if strings.HasPrefix(name, constant.SignaturePrefix) {
			signature = content

			break
		}
	}

	if signature == nil {
		t.Fatalf(
			"signature segment missing .SIGN.RSA.* file, got files: %v",
			slices.Collect(maps.Keys(signatureFiles)),
		)
	}

	controlSegment := signedSegments[1]
	hash := sha1.Sum(controlSegment)
	errors.PanicOnError(
		rsa.VerifyPKCS1v15(
			&privateKey.PublicKey,
			crypto.SHA1,
			hash[:],
			signature,
		),
	)
	controlFiles := package_server_tester.ReadTarGz(signedSegments[1])

	if _, okay := controlFiles[constant.MetadataFile]; !okay {
		t.Errorf("control segment missing .PKGINFO")
	}

	archiveFiles := package_server_tester.ReadTarGz(signedSegments[2])

	if _, okay := archiveFiles["usr/bin/test-binary"]; !okay {
		t.Errorf("data segment missing binary")
	}
}
