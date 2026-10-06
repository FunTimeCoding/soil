package scope_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func HelperWithoutCleanup(
	t *testing.T,
	directory string,
) {
	t.Helper()
	testutil.WriteFile(
		t,
		directory,
		"helper/server.go",
		"package helper\n\nimport (\n\t\"net/http\"\n\t\"testing\"\n)\n\ntype Server struct {\n\tresponse *http.Response\n}\n\nfunc (s *Server) Close() {\n\t_ = s.response.Body.Close()\n}\n\nfunc (s *Server) Ping() {}\n\nfunc NewServer(t *testing.T) *Server {\n\tr, _ := http.Get(\"http://example.com\")\n\tresult := &Server{response: r}\n\tt.Helper()\n\n\treturn result\n}\n",
	)
}
