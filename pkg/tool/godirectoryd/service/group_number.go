package service

import (
	"github.com/funtimecoding/soil/pkg/directory/types/search_record"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"strconv"
)

func groupNumber(record *search_record.Record) int {
	value, e := strconv.Atoi(first(record, constant.GroupNumberAttribute))

	if e != nil {
		return 0
	}

	return value
}
