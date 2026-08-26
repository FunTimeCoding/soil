package goalpine

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/tool/goalpine/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/client"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.String(argumentConstant.Name, "", "Filter to one package name")
	a.Parse(version, gitHash, buildDate)
	c := client.NewEnvironment()
	fmt.Println(c.Packages(a.GetString(argumentConstant.Name)))
}
