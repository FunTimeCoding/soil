package option

import "os"

func New() *Analyze {
	return &Analyze{Notice: os.Stderr}
}
