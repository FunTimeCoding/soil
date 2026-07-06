package debian

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/debian/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
	"os"
)

func PrepareNetBoot(
	inputArchive string,
	outputArchive string,
	configurationPath string,
	temporaryDirectory string,
) {
	system.ExtractTarZip(inputArchive, temporaryDirectory)
	errors.PanicOnError(
		os.Rename(
			configurationPath,
			join.Absolute(temporaryDirectory, constant.PreseedConfiguration),
		),
	)
	system.CreateTarZip(temporaryDirectory, outputArchive)
}
