package check

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/linux"
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
		fmt.Println("Linux")
		fmt.Printf("Cores: %d\n", runtime.NumCPU())
		fmt.Printf("Failed: %s\n", Execute(command.Failed()))
		// TODO: Load average > CPU cores check
		diskFull()

		if run.CommandExists(linux.Jc) {
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
					fmt.Printf("Found port: %s\n", n.LocalPort)
				}

				slices.Sort(ports)
				slices.Sort(found)

				if !slices.Equal(ports, found) {
					fmt.Printf(
						"Expect ports: %s\n",
						join.Comma(ports),
					)
					fmt.Printf(
						"Found ports: %s\n",
						join.Comma(found),
					)
				}
			}
		} else {
			fmt.Println("jc not found")
		}
	case constant.Darwin:
		fmt.Println("Darwin")
	}
}
