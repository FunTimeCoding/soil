package debian

import "fmt"

func RenderUnit(executableName string) string {
	return fmt.Sprintf(
		"[Unit]\nDescription=%s stub description\nAfter=network.target\n\n[Service]\nType=simple\nExecStart=/usr/local/bin/%s\n\n[Install]\nWantedBy=multi-user.target\n",
		executableName,
		executableName,
	)
}
