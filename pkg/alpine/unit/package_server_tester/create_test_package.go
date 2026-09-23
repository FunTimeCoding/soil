package package_server_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/alpine/packager"
	constant1 "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func CreateTestPackage(
	directory string,
	outputPath string,
) {
	scriptPath := filepath.Join(directory, "test-binary")
	scriptContent := "#!/bin/sh\necho test\n"
	system.WriteFile(scriptPath, []byte(scriptContent), 0755)
	p := packager.NewCustom(
		packager.WithExecutablePath(scriptPath),
		packager.WithExecutableName("test-binary"),
		packager.WithPackageVersion(constant1.DefaultVersion),
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
	archiveHash := CreateArchive(p.WorkDirectory, p.ArchiveDirectory)
	pkginfo := fmt.Sprintf(
		"pkgname = test-binary\npkgver = 1.0.0\narch = x86_64\ndatahash = %s\n",
		archiveHash,
	)
	system.WriteFile(
		filepath.Join(p.ControlDirectory, constant.MetadataFile),
		[]byte(pkginfo),
		0644,
	)
	CreateControlTar(p.WorkDirectory, p.ControlDirectory)
	ConcatenateFiles(
		outputPath,
		filepath.Join(p.WorkDirectory, constant.ControlFile),
		filepath.Join(p.WorkDirectory, constant.ArchiveFile),
	)
}
