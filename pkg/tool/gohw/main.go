package gohw

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/netbox"
	netboxConstant "github.com/funtimecoding/soil/pkg/netbox/constant"
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
	d := n.MustDeviceByName(hostname)
	f := consoleConstant.ExtendedColorFormat.Copy().Tag(
		consoleConstant.TagIdentifier,
	)

	if d == nil {
		r := n.MustDeviceRoleByName("default")
		t := n.MustDeviceTypeByName("default")
		s := n.MustSiteByName("default")
		e := n.MustTenantByName("default")
		console.Format("Hostname: %s\n", hostname)
		console.Format("Role: %s\n", r.Format(f))
		console.Format("Type: %s\n", t.Format(f))
		console.Format("Site: %s\n", s.Format(f))
		console.Format("Tenant: %s\n", e.Format(f))
		ip, mask, g := primaryInterface()

		if g != nil {
			log.Panicf("interface fail: %v", g)
		}

		ones, _ := mask.Size()
		console.Format("Primary IP: %s/%d\n", ip.String(), ones)

		if false {
			n.MustCreateDevice(hostname, r, []string{}, t, s, e)
			d = n.MustDeviceByName(hostname)

			if d == nil {
				panic("retrieve fail")
			}
		}
	}

	if false {
		if d.PrimaryAddress == "" {
			ip, mask, e := primaryInterface()

			if e != nil {
				log.Panicf("interface fail: %v", e)
			}

			objectType := netboxConstant.DeviceAddress
			objectIdentifier := int64(d.Identifier)
			n.MustCreateInternet(objectType, objectIdentifier, ip, mask)
		}
	}
}
