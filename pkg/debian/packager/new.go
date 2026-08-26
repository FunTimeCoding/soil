package packager

import (
	"github.com/funtimecoding/soil/pkg/debian"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
	"path/filepath"
)

func New(
	pathToExecutable string,
	version string,
	name string,
	mail string,
) *Packager {
	architecture := constant.AMD64
	executableName := filepath.Base(pathToExecutable)
	packageName := debian.PackageVersion(
		executableName,
		version,
		1,
		architecture,
	)
	root := join.Absolute(system.WorkDirectory(), packageName)
	configuration := configurationRoot(root)

	return &Packager{
		Architecture:      architecture,
		ExecutablePath:    pathToExecutable,
		ExecutableName:    executableName,
		PackageName:       packageName,
		PackageVersion:    version,
		MaintainerName:    name,
		MaintainerMail:    mail,
		Root:              root,
		ConfigurationRoot: configuration,
		UnitRoot:          unitRoot(root),
		ControlFile: join.Absolute(
			configuration,
			constant.DebianControlFile,
		),
		BinaryRoot: binaryRoot(root),
	}
}
