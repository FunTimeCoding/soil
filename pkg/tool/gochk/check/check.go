package check

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/split"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"runtime"
	"slices"
)

func Check(port string) {
	switch runtime.GOOS {
	case systemConstant.Linux:
		console.Line("Linux")
		console.Format("Cores: %d\n", runtime.NumCPU())
		console.Format(
			"Failed: %s\n",
			Execute(systemConstant.SystemdFailedCommand),
		)
		diskFull()

		if run.CommandExists(systemConstant.Jc) {
			if port != "" {
				ports := split.Comma(port)
				var found []string

				for _, n := range Netstat(false) {
					if n.NetworkProtocol != "ipv4" {
						continue
					}

					if !slices.Contains(ports, n.LocalPort) {
						continue
					}

					found = append(found, n.LocalPort)
					console.Format("Found port: %s\n", n.LocalPort)
				}

				slices.Sort(ports)
				slices.Sort(found)

				if !slices.Equal(ports, found) {
					console.Format("Expect ports: %s\n", join.Comma(ports))
					console.Format("Found ports: %s\n", join.Comma(found))
				}
			}
		} else {
			console.Line("jc not found")
		}
	case systemConstant.Darwin:
		console.Line("Darwin")
	}
}
