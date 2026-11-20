package unit

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/unit/scan_tester"
	"testing"
)

func TestErrorResponseOnCaptured503IsClean(t *testing.T) {
	v := virtual_file_system.New()
	scan_tester.WriteServedSpec(v, scan_tester.ErrorResponseSpec("503"))
	scan_tester.New(t).NoConcern(
		constant.ErrorResponseOnNon500Key,
		scan_tester.New(t).One(v, false),
	)
}

func TestErrorResponseOnA404IsFlagged(t *testing.T) {
	v := virtual_file_system.New()
	scan_tester.WriteServedSpec(v, scan_tester.ErrorResponseSpec("404"))
	scan_tester.New(t).Concern(
		constant.ErrorResponseOnNon500Key,
		scan_tester.New(t).One(v, false),
	)
}
