package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Status() {
	s := common.Alertmanager().MustStatus()
	console.Format("Status: %+v\n", s)
	console.Format("  Cluster: %+v\n", s.Cluster)
	console.Format("    Status: %s\n", *s.Cluster.Status)
	console.Format("  Configuration: %+v\n", *s.Config.Original)
	console.Format("  Version: %+v\n", s.VersionInfo)
	console.Format("    Branch: %+v\n", *s.VersionInfo.Branch)
	console.Format("    Version: %+v\n", *s.VersionInfo.Version)
	console.Format("    Revision: %+v\n", *s.VersionInfo.Revision)
	console.Format("    GoVersion: %+v\n", *s.VersionInfo.GoVersion)
	console.Format("    BuildDate: %+v\n", *s.VersionInfo.BuildDate)
}
