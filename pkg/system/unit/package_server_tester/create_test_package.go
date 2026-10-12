package package_server_tester

import (
	"fmt"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/linux/alpine/packager"
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
		packager.WithPackageVersion(soil.DefaultVersion),
		packager.WithWorkDirectory(filepath.Join(directory, "workspace")),
		packager.WithControlDirectory(
			filepath.Join(
				directory,
				"workspace",
				constant.AlpineControlDirectory,
			),
		),
		packager.WithArchiveDirectory(
			filepath.Join(directory, "workspace", "data"),
		),
		packager.WithOutputFile(outputPath),
	)
	system.MakeDirectory(p.ControlDirectory)
	system.MakeDirectory(p.ArchiveDirectory)
	binDirectory := filepath.Join(
		p.ArchiveDirectory,
		constant.Resources,
		constant.Binary,
	)
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
		filepath.Join(p.ControlDirectory, constant.AlpineMetadataFile),
		[]byte(pkginfo),
		0644,
	)
	CreateControlTar(p.WorkDirectory, p.ControlDirectory)
	ConcatenateFiles(
		outputPath,
		filepath.Join(p.WorkDirectory, constant.AlpineControlFile),
		filepath.Join(p.WorkDirectory, constant.AlpineArchiveFile),
	)
}
