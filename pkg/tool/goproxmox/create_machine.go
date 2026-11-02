package goproxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
	"github.com/spf13/cobra"
)

func createMachine(c *command_context.Context) *cobra.Command {
	f := &machineFlags{}
	result := &cobra.Command{
		Use:   "create-machine",
		Short: "Create a virtual machine",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			runCreateMachine(c, f)
		},
	}
	result.Flags().StringVar(&f.node, "node", "", "target node")
	result.Flags().StringVar(&f.name, "name", "", "VM name")
	result.Flags().StringVar(
		&f.diskImport,
		"disk-import",
		"",
		"import-from path",
	)
	result.Flags().IntVar(&f.diskSize, "disk-size", 0, "disk size in GiB")
	result.Flags().StringVar(
		&f.diskStorage,
		"disk-storage",
		"",
		"storage backend",
	)
	result.Flags().StringVar(&f.cdrom, "cdrom", "", "ISO volume")
	result.Flags().StringVar(&f.bridge, "bridge", "", "network bridge")
	result.Flags().StringVar(&f.cpuType, "cpu-type", "", "CPU type")
	result.Flags().StringVar(&f.operatingSystemType, "os-type", "", "OS type")
	result.Flags().StringVar(&f.ciUser, "ci-user", "", "cloud-init user")
	result.Flags().StringVar(
		&f.ciPassword,
		"ci-password",
		"",
		"cloud-init password",
	)
	result.Flags().StringVar(&f.sshKeys, "ssh-keys", "", "SSH public keys")
	result.Flags().StringVar(
		&f.ipConfiguration,
		"ip-config",
		"",
		"cloud-init IP config",
	)
	result.Flags().StringVar(
		&f.searchDomain,
		"search-domain",
		"",
		"DNS search domain",
	)
	result.Flags().IntVar(&f.cores, "cores", 0, "CPU cores")
	result.Flags().IntVar(&f.memory, "memory", 0, "memory in MiB")
	result.Flags().StringVar(&f.tags, "tags", "", "semicolon-separated tags")
	result.Flags().StringVar(&f.extras, "extras", "", "extra key=value pairs")
	result.Flags().BoolVar(&f.start, "start", false, "start after creation")
	errors.PanicOnError(result.MarkFlagRequired("node"))
	errors.PanicOnError(result.MarkFlagRequired("name"))

	return result
}
