package goproxmox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/spf13/cobra"
)

func createMachine(c *command_context.Context) *cobra.Command {
	var node string
	var name string
	var diskImport string
	var diskSize int
	var diskStorage string
	var cdrom string
	var bridge string
	var cpuType string
	var osType string
	var ciUser string
	var ciPassword string
	var sshKeys string
	var ipConfiguration string
	var searchDomain string
	var cores int
	var memory int
	var tags string
	var extras string
	var start bool
	result := &cobra.Command{
		Use:   "create-machine",
		Short: "Create a virtual machine",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			body := client.CreateMachineJSONRequestBody{
				Node: node,
				Name: name,
			}

			if diskImport != "" {
				body.DiskImport = &diskImport
			}

			if diskSize > 0 {
				body.DiskSize = &diskSize
			}

			if diskStorage != "" {
				body.DiskStorage = &diskStorage
			}

			if cdrom != "" {
				body.Cdrom = &cdrom
			}

			if bridge != "" {
				body.Bridge = &bridge
			}

			if cpuType != "" {
				body.CpuType = &cpuType
			}

			if osType != "" {
				body.OsType = &osType
			}

			if ciUser != "" {
				body.CiUser = &ciUser
			}

			if ciPassword != "" {
				body.CiPassword = &ciPassword
			}

			if sshKeys != "" {
				body.SshKeys = &sshKeys
			}

			if ipConfiguration != "" {
				body.IpConfig = &ipConfiguration
			}

			if searchDomain != "" {
				body.SearchDomain = &searchDomain
			}

			if cores > 0 {
				body.Cores = &cores
			}

			if memory > 0 {
				body.Memory = &memory
			}

			if tags != "" {
				body.Tags = &tags
			}

			if extras != "" {
				body.Extras = &extras
			}

			if start {
				body.Start = &start
			}

			fmt.Println(c.Client().CreateMachine(body))
		},
	}
	result.Flags().StringVar(&node, "node", "", "target node")
	result.Flags().StringVar(&name, "name", "", "VM name")
	result.Flags().StringVar(
		&diskImport,
		"disk-import",
		"",
		"import-from path",
	)
	result.Flags().IntVar(&diskSize, "disk-size", 0, "disk size in GiB")
	result.Flags().StringVar(
		&diskStorage,
		"disk-storage",
		"",
		"storage backend",
	)
	result.Flags().StringVar(&cdrom, "cdrom", "", "ISO volume")
	result.Flags().StringVar(&bridge, "bridge", "", "network bridge")
	result.Flags().StringVar(&cpuType, "cpu-type", "", "CPU type")
	result.Flags().StringVar(&osType, "os-type", "", "OS type")
	result.Flags().StringVar(&ciUser, "ci-user", "", "cloud-init user")
	result.Flags().StringVar(
		&ciPassword,
		"ci-password",
		"",
		"cloud-init password",
	)
	result.Flags().StringVar(&sshKeys, "ssh-keys", "", "SSH public keys")
	result.Flags().StringVar(
		&ipConfiguration,
		"ip-config",
		"",
		"cloud-init IP config",
	)
	result.Flags().StringVar(
		&searchDomain,
		"search-domain",
		"",
		"DNS search domain",
	)
	result.Flags().IntVar(&cores, "cores", 0, "CPU cores")
	result.Flags().IntVar(&memory, "memory", 0, "memory in MiB")
	result.Flags().StringVar(&tags, "tags", "", "semicolon-separated tags")
	result.Flags().StringVar(&extras, "extras", "", "extra key=value pairs")
	result.Flags().BoolVar(&start, "start", false, "start after creation")
	errors.PanicOnError(result.MarkFlagRequired("node"))
	errors.PanicOnError(result.MarkFlagRequired("name"))

	return result
}
