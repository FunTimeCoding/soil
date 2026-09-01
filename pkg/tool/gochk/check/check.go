package check

import (
	"github.com/funtimecoding/soil/pkg/console"
	linuxConstant "github.com/funtimecoding/soil/pkg/linux/constant"
	"github.com/funtimecoding/soil/pkg/linux/systemd/command"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"runtime"
	"slices"
)

func Check(port string) {
	switch runtime.GOOS {
	case constant.Linux:
		console.Line("Linux")
		console.Format("Cores: %d\n", runtime.NumCPU())
		console.Format("Failed: %s\n", Execute(command.Failed()))
		// TODO: Load average > CPU cores check
		diskFull()

		if run.CommandExists(linuxConstant.Jc) {
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
	case constant.Darwin:
		console.Line("Darwin")
	}
}
