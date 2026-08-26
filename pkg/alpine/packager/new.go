package packager

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"os"
	"path/filepath"
)

func New(
	pathToExecutable string,
	version string,
) *Packager {
	executableName := filepath.Base(pathToExecutable)
	packageVersion := fmt.Sprintf("%s-r%s", version, constant.Release)
	packageName := fmt.Sprintf("%s-%s.apk", executableName, packageVersion)
	work := filepath.Join(
		os.TempDir(),
		fmt.Sprintf("gopackageapk-%s", executableName),
	)

	return &Packager{
		ExecutablePath:   pathToExecutable,
		ExecutableName:   executableName,
		PackageVersion:   packageVersion,
		PackageName:      packageName,
		WorkDirectory:    work,
		ControlDirectory: filepath.Join(work, "control"),
		ArchiveDirectory: filepath.Join(work, "data"),
		OutputFile:       packageName,
	}
}
