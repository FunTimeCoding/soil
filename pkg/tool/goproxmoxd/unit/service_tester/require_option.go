package service_tester

import (
	"github.com/luthermonson/go-proxmox"
	"testing"
)

func RequireOption(
	t *testing.T,
	options []proxmox.VirtualMachineOption,
	name string,
) any {
	t.Helper()
	v, okay := FindOption(options, name)

	if !okay {
		t.Fatalf("option %q not found", name)
	}

	return v
}
