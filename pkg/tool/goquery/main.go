package goquery

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goquery/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/spf13/cobra"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	host := environment.Required(constant.HostEnvironment)
	port := environment.RequiredInteger(constant.PortEnvironment)
	c, e := client.NewClient(locator.New(host).Port(port).Insecure().String())
	errors.PanicOnError(e)
	o := &cobra.Command{
		Use:     constant.Identity.Usage(),
		Short:   constant.Identity.Description(),
		Version: argument.CobraVersion(version, gitHash, buildDate),
	}
	o.AddCommand(search(c))
	o.AddCommand(list(c))
	o.AddCommand(index(c))
	o.AddCommand(addCollection(c))
	o.AddCommand(deleteCollection(c))
	o.AddCommand(embed(c))
	o.AddCommand(addContext(c))
	o.AddCommand(listContexts(c))
	o.AddCommand(listTags(c))
	o.AddCommand(listMetadata(c))
	o.AddCommand(removeContext(c))
	o.AddCommand(status(c))
	errors.PanicOnError(o.Execute())
}
