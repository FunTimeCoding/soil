package goalertlogd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/option"
	"path/filepath"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.String(
		argument.Path,
		filepath.Join(
			constant.Identity.StorageDirectory(false),
			join.Empty(constant.Identity.Name(), library.BoltExtension),
		),
		"Database path",
	)
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Address = a.Address()
	o.DatabasePath = a.GetString(argument.Path)
	o.Version = version
	Run(o, r)
}
