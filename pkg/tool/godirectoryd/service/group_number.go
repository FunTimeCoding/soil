package service

import (
	"github.com/funtimecoding/soil/pkg/directory"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"strconv"
)

func groupNumber(record *directory.Record) int {
	value, e := strconv.Atoi(first(record, constant.GroupNumberAttribute))

	if e != nil {
		return 0
	}

	return value
}
