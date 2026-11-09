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

func Main() {
	r := reporter.New(constant.Identity.Name())
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Parse()
	hostname := host.StripDomain(system.Hostname())
	n := netbox.NewEnvironment()
	defer n.Close()

	if n.MustDevice(hostname) != nil {
		return
	}

	f := consoleConstant.ExtendedColorFormat.Copy().Tag(
		consoleConstant.TagIdentifier,
	)
	console.Format("Hostname: %s\n", hostname)
	console.Format("Role: %s\n", n.MustDeviceRole("default").Format(f))
	console.Format("Type: %s\n", n.MustDeviceType("default").Format(f))
	console.Format("Site: %s\n", n.MustSite("default").Format(f))
	console.Format("Tenant: %s\n", n.MustTenant("default").Format(f))
	ip, mask, e := primaryInterface()

	if e != nil {
		log.Panicf("interface fail: %v", e)
	}

	ones, _ := mask.Size()
	console.Format("Primary IP: %s/%d\n", ip.String(), ones)
}
