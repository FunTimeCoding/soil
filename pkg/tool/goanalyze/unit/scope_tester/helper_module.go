package scope_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
)

func (o *Tester) HelperModule(
	directory string,
	module string,
) {
	o.t.Helper()
	system_tester.New(o.t).Write(
		directory,
		"helper/server.go",
		"package helper\n\nimport (\n\t\"net/http\"\n\t\"testing\"\n)\n\ntype Server struct {\n\tresponse *http.Response\n}\n\nfunc (s *Server) Close() {\n\t_ = s.response.Body.Close()\n}\n\nfunc (s *Server) Ping() {}\n\nfunc NewServer(t *testing.T) *Server {\n\tr, _ := http.Get(\"http://example.com\")\n\tresult := &Server{response: r}\n\tt.Cleanup(result.Close)\n\n\treturn result\n}\n",
	)
	system_tester.New(o.t).Write(
		directory,
		"helper/leak.go",
		"package helper\n\nimport \"net/http\"\n\nfunc Leak() int {\n\tr, _ := http.Get(\"http://example.com\")\n\n\treturn r.StatusCode\n}\n",
	)
	testutil.WriteModFile(o.t, directory, module)
}
