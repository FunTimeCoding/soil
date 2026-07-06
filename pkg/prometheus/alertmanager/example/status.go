package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Status() {
	s := common.Alertmanager().MustStatus()
	fmt.Printf("Status: %+v\n", s)
	fmt.Printf("  Cluster: %+v\n", s.Cluster)
	fmt.Printf("    Status: %s\n", *s.Cluster.Status)
	fmt.Printf("  Configuration: %+v\n", *s.Config.Original)
	fmt.Printf("  Version: %+v\n", s.VersionInfo)
	fmt.Printf("    Branch: %+v\n", *s.VersionInfo.Branch)
	fmt.Printf("    Version: %+v\n", *s.VersionInfo.Version)
	fmt.Printf("    Revision: %+v\n", *s.VersionInfo.Revision)
	fmt.Printf("    GoVersion: %+v\n", *s.VersionInfo.GoVersion)
	fmt.Printf("    BuildDate: %+v\n", *s.VersionInfo.BuildDate)
}
