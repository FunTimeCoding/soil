package god8y

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/text/dictionary/check/duplicate"
	"github.com/funtimecoding/soil/pkg/text/dictionary/check/missing"
	"github.com/funtimecoding/soil/pkg/text/dictionary/check/order"
	"github.com/funtimecoding/soil/pkg/tool/god8y/constant"
	"os"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()

	if len(os.Args) >= 3 && os.Args[1] == "merge" {
		merge(os.Args[2:])

		return
	}

	a := argument.NewInstance(constant.Identity)
	a.Parse(version, gitHash, buildDate)
	order.Run()
	missing.Run()
	duplicate.Run()
}
