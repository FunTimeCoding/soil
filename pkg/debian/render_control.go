package debian

import "fmt"

func RenderControl(
	executableName string,
	architecture string,
	packageVersion string,
	maintainerName string,
	maintainerMail string,
) string {
	return fmt.Sprintf(
		"Package: %s\nVersion: %s\nArchitecture: %s\nMaintainer: %s <%s>\nDescription: Short stub description.\n Long stub description.\n",
		executableName,
		packageVersion,
		architecture,
		maintainerName,
		maintainerMail,
	)
}
