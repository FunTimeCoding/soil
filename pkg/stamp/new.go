package stamp

import "runtime/debug"

func New() *Stamp {
	i, okay := debug.ReadBuildInfo()

	if !okay {
		return Read(nil)
	}

	return Read(i)
}
