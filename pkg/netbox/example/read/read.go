package read

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netbox"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
)

func Read() {
	n := netbox.NewEnvironment()
	f := constant.Format
	readTenant(n, f)
	readDCIM(n, f)
	readIPAM(n, f)
	readVirtual(n, f)
	readUser(n, f)
	readExtra(n, f)
	readWireless(n, f)
	readTunnel(n, f)
	readCore(n, f)
	// TODO: DataSource: Requires local, git or S3 source
	//  Is this used to import entities?
	for _, s := range n.MustSources() {
		fmt.Printf("DataSource: %s\n", s.Format(f))
	}

	// TODO: ExportTemplate, this optionally depends on DataSource
	//  Whats a good use case?
	for _, t := range n.MustExportTemplates() {
		fmt.Printf("ExportTemplate: %s\n", t.Format(f))
	}

	// TODO: CustomField, CustomFieldChoice, CustomLink
	//  Whats a good use case?
	for _, i := range n.MustCustomFields() {
		fmt.Printf("CustomField: %s\n", i.Format(f))
	}

	for _, c := range n.MustCustomFieldChoices() {
		fmt.Printf("CustomFieldChoice: %s\n", c.Format(f))
	}

	for _, l := range n.MustCustomLinks() {
		fmt.Printf("CustomLink: %s\n", l.Format(f))
	}

	// TODO: Circuits, CircuitTypes, VirtualCircuits, VirtualCircuitTypes, CircuitGroups, Providers, ProviderAccounts, ProviderNetworks
	// TODO: Create CircuitTermination fails: A circuit termination must attach to a terminating object.
	// TODO: Create VirtualCircuitTermination requires an interface somewhere
	// TODO: Create GroupAssignment fails: null value in column "member_id" of relation "circuits_circuitgroupassignment" violates not-null constraint
}
