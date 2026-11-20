package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/flag"
)

func Flags(v []flag.Flag) []*server.Flag {
	result := make([]*server.Flag, len(v))

	for i := range v {
		result[i] = Flag(&v[i])
	}

	return result
}
