package gowait

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
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
	a.String(argumentConstant.File, "", "File to wait for")
	a.String(argumentConstant.Process, "", "Process to wait for")
	a.String(argumentConstant.Locator, "", "Locator to wait for")
	a.String(argumentConstant.Contains, "", "String for locator")
	a.Duration(argumentConstant.Timeout, 3*time.Minute, "")
	a.Boolean(argumentConstant.Verbose, false, "Verbose output")
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.File = a.GetString(argumentConstant.File)
	o.Process = a.GetString(argumentConstant.Process)
	o.Locator = a.GetString(argumentConstant.Locator)
	o.Contains = a.GetString(argumentConstant.Contains)
	o.Timeout = a.GetDuration(argumentConstant.Timeout)
	o.Verbose = a.GetBoolean(argumentConstant.Verbose)
	wait.Run(o)
}
