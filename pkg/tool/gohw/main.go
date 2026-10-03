package gohw

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/netbox"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gohw/constant"
	"github.com/funtimecoding/soil/pkg/web/host"
	"log"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version)
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Parse(version, gitHash, buildDate)
	hostname := host.StripDomain(system.Hostname())
	n := netbox.NewEnvironment()

	if n.MustDeviceByName(hostname) != nil {
		return
	}

	f := consoleConstant.ExtendedColorFormat.Copy().Tag(
		consoleConstant.TagIdentifier,
	)
	console.Format("Hostname: %s\n", hostname)
	console.Format("Role: %s\n", n.MustDeviceRoleByName("default").Format(f))
	console.Format("Type: %s\n", n.MustDeviceTypeByName("default").Format(f))
	console.Format("Site: %s\n", n.MustSiteByName("default").Format(f))
	console.Format("Tenant: %s\n", n.MustTenantByName("default").Format(f))
	ip, mask, e := primaryInterface()

	if e != nil {
		log.Panicf("interface fail: %v", e)
	}

	ones, _ := mask.Size()
	console.Format("Primary IP: %s/%d\n", ip.String(), ones)
}
