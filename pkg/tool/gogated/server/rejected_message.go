package server

import "github.com/funtimecoding/soil/pkg/tool/gogated/constant"

func rejectedMessage(source string) string {
	if source == constant.SourceDirectory {
		return constant.DirectoryRejectedMessage
	}

	return constant.LocalRejectedMessage
}
