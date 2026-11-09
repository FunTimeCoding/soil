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
	assertNoConcern(t, scanOne(t, v, false), constant.ErrorResponseOnNon500Key)
}

func TestErrorResponseOnA404IsFlagged(t *testing.T) {
	v := virtual_file_system.New()
	scan_tester.WriteServedSpec(v, scan_tester.ErrorResponseSpec("404"))
	assertConcern(t, scanOne(t, v, false), constant.ErrorResponseOnNon500Key)
}
