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
	scan_tester.NoConcern(
		t,
		constant.ErrorResponseOnNon500Key,
		scan_tester.One(t, v, false),
	)
}

func TestErrorResponseOnA404IsFlagged(t *testing.T) {
	v := virtual_file_system.New()
	scan_tester.WriteServedSpec(v, scan_tester.ErrorResponseSpec("404"))
	scan_tester.Concern(
		t,
		constant.ErrorResponseOnNon500Key,
		scan_tester.One(t, v, false),
	)
}
