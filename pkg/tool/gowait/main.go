package gowait

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/tool/gowait/constant"
	"github.com/funtimecoding/soil/pkg/tool/gowait/wait"
	"github.com/funtimecoding/soil/pkg/tool/gowait/wait/option"
	"time"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.String(argument.File, "", "File to wait for")
	a.String(argument.Process, "", "Process to wait for")
	a.String(argument.Locator, "", "Locator to wait for")
	a.String(argument.Contains, "", "String for locator")
	a.Duration(argument.Timeout, 3*time.Minute, "")
	a.Boolean(argument.Verbose, false, "Verbose output")
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.File = a.GetString(argument.File)
	o.Process = a.GetString(argument.Process)
	o.Locator = a.GetString(argument.Locator)
	o.Contains = a.GetString(argument.Contains)
	o.Timeout = a.GetDuration(argument.Timeout)
	o.Verbose = a.GetBoolean(argument.Verbose)
	wait.Run(o)
}
