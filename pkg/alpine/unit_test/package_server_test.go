package unit_test

import (
	"archive/tar"
	"compress/gzip"
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/alpine/package_server"
	"github.com/funtimecoding/soil/pkg/alpine/packager"
	alpineTest "github.com/funtimecoding/soil/pkg/alpine/testing"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSign(t *testing.T) {
	d := filepath.Join("../../..", "tmp", "alpine-signing-test")
	system.MakeDirectory(d)
	unsignedPath := filepath.Join(d, "test-unsigned.apk")
	createTestPackage(d, unsignedPath)
	privateKey := alpineTest.GenerateRSAKey(2048)
	unsignedSegments := alpineTest.SplitArchive(unsignedPath)

	if len(unsignedSegments) != 2 {
		t.Fatalf(
			"unsigned package should have 2 segments, got %d",
			len(unsignedSegments),
		)
	}

	package_server.SignPackageWithKey(unsignedPath, privateKey, "test.rsa")
	signedSegments := alpineTest.SplitArchive(unsignedPath)

	if len(signedSegments) != 3 {
		t.Fatalf(
			"signed package should have 3 segments, got %d",
			len(signedSegments),
		)
	}

	signatureSegment := signedSegments[0]
	signatureFiles := alpineTest.ReadTarGz(signatureSegment)
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
			getKeys(signatureFiles),
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
	controlFiles := alpineTest.ReadTarGz(signedSegments[1])

	if _, okay := controlFiles[constant.MetadataFile]; !okay {
		t.Errorf("control segment missing .PKGINFO")
	}

	archiveFiles := alpineTest.ReadTarGz(signedSegments[2])

	if _, okay := archiveFiles["usr/bin/test-binary"]; !okay {
		t.Errorf("data segment missing binary")
	}
}

func createTestPackage(
	directory string,
	outputPath string,
) {
	scriptPath := filepath.Join(directory, "test-binary")
	scriptContent := "#!/bin/sh\necho test\n"
	system.WriteFile(scriptPath, []byte(scriptContent), 0755)
	p := packager.NewCustom(
		packager.WithExecutablePath(scriptPath),
		packager.WithExecutableName("test-binary"),
		packager.WithPackageVersion(library.DefaultVersion),
		packager.WithWorkDirectory(filepath.Join(directory, "workspace")),
		packager.WithControlDirectory(
			filepath.Join(directory, "workspace", "control"),
		),
		packager.WithArchiveDirectory(
			filepath.Join(directory, "workspace", "data"),
		),
		packager.WithOutputFile(outputPath),
	)
	system.MakeDirectory(p.ControlDirectory)
	system.MakeDirectory(p.ArchiveDirectory)
	binDirectory := filepath.Join(p.ArchiveDirectory, "usr", "bin")
	system.MakeDirectory(binDirectory)
	system.WriteFile(
		filepath.Join(binDirectory, "test-binary"),
		system.ReadBytesUnsafe(scriptPath),
		0755,
	)
	archiveHash := createArchive(p.WorkDirectory, p.ArchiveDirectory)
	pkginfo := fmt.Sprintf(
		"pkgname = test-binary\npkgver = 1.0.0\narch = x86_64\ndatahash = %s\n",
		archiveHash,
	)
	system.WriteFile(
		filepath.Join(p.ControlDirectory, constant.MetadataFile),
		[]byte(pkginfo),
		0644,
	)
	createControlTar(p.WorkDirectory, p.ControlDirectory)
	concatenateFiles(
		outputPath,
		filepath.Join(p.WorkDirectory, constant.ControlFile),
		filepath.Join(p.WorkDirectory, constant.ArchiveFile),
	)
}

func concatenateFiles(
	output string,
	files ...string,
) {
	f, e := os.Create(output)
	errors.PanicOnError(e)
	defer errors.PanicClose(f)

	for _, i := range files {
		b, g := os.ReadFile(i)
		errors.PanicOnError(g)
		_, h := f.Write(b)
		errors.PanicOnError(h)
	}
}

func getKeys(m map[string][]byte) []string {
	result := make([]string, 0, len(m))

	for k := range m {
		result = append(result, k)
	}

	return result
}

func createControlTar(
	workDirectory string,
	controlDirectory string,
) {
	f := system.Create(filepath.Join(workDirectory, constant.ControlFile))
	defer errors.PanicClose(f)
	z := gzip.NewWriter(f)
	defer errors.PanicClose(z)
	t := tar.NewWriter(z)
	p := system.Open(filepath.Join(controlDirectory, constant.MetadataFile))
	defer errors.PanicClose(p)
	errors.PanicOnError(
		t.WriteHeader(
			&tar.Header{
				Name: constant.MetadataFile,
				Size: system.FileStat(p).Size(),
				Mode: 0644,
			},
		),
	)
	system.Copy(p, t)
	errors.PanicFlush(t)
}

func createArchive(
	workDirectory string,
	archiveDirectory string,
) string {
	path := filepath.Join(workDirectory, constant.ArchiveFile)
	f := system.Create(path)
	defer errors.PanicClose(f)
	z := gzip.NewWriter(f)
	defer errors.PanicClose(z)
	t := tar.NewWriter(z)
	defer errors.PanicClose(t)
	errors.PanicOnError(
		filepath.Walk(
			archiveDirectory,
			func(
				path string,
				i os.FileInfo,
				e error,
			) error {
				errors.PanicOnError(e)

				if path == archiveDirectory {
					return nil
				}

				relPath := system.RelativePath(archiveDirectory, path)
				h := system.TarHeader(i, "")
				h.Name = filepath.ToSlash(relPath)
				h.Format = tar.FormatPAX
				system.TarWriteHeader(t, h)

				if i.Mode().IsRegular() {
					file := system.Open(path)
					defer errors.PanicClose(file)
					system.Copy(file, t)
				}

				return nil
			},
		),
	)
	o := system.Open(path)
	defer errors.PanicClose(o)
	h := sha256.New()
	system.Copy(o, h)

	return hex.EncodeToString(h.Sum(nil))
}
