package format

import "fmt"

func MiB(bytes int64) string {
	mib := float64(bytes) / 1024 / 1024

	if mib >= 1024 {
		return fmt.Sprintf("%.1f GiB", mib/1024)
	}

	return fmt.Sprintf("%.0f MiB", mib)
}
