package crap

import (
	"github.com/funtimecoding/soil/pkg/console/table"
	"github.com/funtimecoding/soil/pkg/crap/baseline"
	"github.com/funtimecoding/soil/pkg/crap/constant"
)

func newScoreTable(b *baseline.Baseline) *table.Table {
	headers := []string{
		constant.ColumnMark,
		constant.ColumnScore,
		constant.ColumnCC,
		constant.ColumnCoverage,
	}
	right := []int{1, 2, 3}

	if b != nil {
		headers = append(headers, constant.ColumnDelta)
		right = append(right, 4)
	}

	return table.New(
		append(headers, constant.ColumnFunction, constant.ColumnLocation)...,
	).Right(right...)
}
