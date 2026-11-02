package constant

var ReleasingMethods = map[string]bool{
	"(*io.PipeReader).CloseWithError": true,
	"(*io.PipeWriter).CloseWithError": true,
}
