package format

import "fmt"

func MilliCores(millis int64) string {
	if millis >= 1000 {
		return fmt.Sprintf("%.1f cores", float64(millis)/1000)
	}

	return fmt.Sprintf("%dm", millis)
}
