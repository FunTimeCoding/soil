package constant

import "regexp"

var CaptureCalls = []string{
	"answerWith",
	"answer",
	"captureFail",
	"captureDetail",
	"CaptureException",
	"CaptureWithContext",
}

var ResponseStatus = regexp.MustCompile(`(\d{3})JSONResponse$`)

const (
	ErrorAnswerKey        = "error_answer"
	ErrorClassSegment     = "/pkg/errors/"
	ErrorClassPredicate   = "Is"
	ErrorsPackage         = "errors"
	AsTypeFunction        = "AsType"
	AsFunction            = "As"
	ResponsePackageSuffix = "/generative/mark/response"
	ResponseFail          = "Fail"
	CaptureFailMethod     = "captureFail"
	ErrorAnswerReference  = "doc/ai/spec/error-handling/README.md, Answering an error"
)
