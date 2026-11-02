package goproxmox

type machineFlags struct {
	node                string
	name                string
	diskImport          string
	diskSize            int
	diskStorage         string
	cdrom               string
	bridge              string
	cpuType             string
	operatingSystemType string
	ciUser              string
	ciPassword          string
	sshKeys             string
	ipConfiguration     string
	searchDomain        string
	cores               int
	memory              int
	tags                string
	extras              string
	start               bool
}
