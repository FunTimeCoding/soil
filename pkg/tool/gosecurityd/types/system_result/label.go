package system_result

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
)

func (f *Finding) Label() string {
	if f.Backports {
		return join.Empty(f.Source, constant.BackportsLabel)
	}

	return f.Source
}
