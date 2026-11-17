package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"strings"
	"testing"
)

func TestIndexedFindReferencesMatchesFullLoad(t *testing.T) {
	o, full := rewrite_tester.IndexedAndFull(t)
	symbols := [][2]string{
		{"NewServer", ""},
		{"Start", "Server"},
		{"Port", "Server"},
		{"Server", ""},
	}

	for _, s := range symbols {
		rewrite_tester.SameResult(
			t,
			func(x *service_tester.Services) (any, any, error) {
				return x.Reference.FindReferences(
					o.Directory,
					"example/alfa",
					s[0],
					s[1],
				)
			},
			o.Services,
			full,
		)
	}

	_, references, e := o.Reference.FindReferences(
		o.Directory,
		"example/alfa",
		"NewServer",
		"",
	)
	assert.FatalOnError(t, e)
	assert.String(
		t,
		`[{"file":"alfa/server.go","line":14,"package":"example/alfa"},{"file":"alfa/server_test.go","line":6,"package":"example/alfa"},{"file":"bravo/external_test.go","line":9,"package":"example/bravo_test"},{"file":"bravo/use.go","line":6,"package":"example/bravo"}]`,
		string(notation.Marshal(references.Locations)),
	)
}

func TestFieldReferencesCrossTestVariant(t *testing.T) {
	o, full := rewrite_tester.IndexedAndFull(t)

	for _, s := range []*service_tester.Services{o.Services, full} {
		_, references, e := s.Reference.FindReferences(
			o.Directory,
			"example/alfa",
			"Port",
			"Server",
		)
		assert.FatalOnError(t, e)
		assert.String(
			t,
			`[{"file":"bravo/use.go","line":9,"package":"example/bravo"},{"file":"charlie/unit/only_test.go","line":9,"package":"example/charlie/unit"},{"file":"echo/pool.go","line":6,"package":"example/echo"},{"file":"echo/pool.go","line":6,"package":"example/echo"}]`,
			string(notation.Marshal(references.Locations)),
		)
	}
}

func TestIndexedFileReferencesMatchesFullLoad(t *testing.T) {
	o, full := rewrite_tester.IndexedAndFull(t)
	rewrite_tester.SameResult(
		t,
		func(x *service_tester.Services) (any, any, error) {
			return x.Reference.FileReferences(
				o.Directory,
				"example/alfa",
				"alfa/server.go",
			)
		},
		o.Services,
		full,
	)
}

func TestIndexedLiteralsReachElisionThroughThirdPackage(t *testing.T) {
	o, full := rewrite_tester.IndexedAndFull(t)
	result := rewrite_tester.SameResult(
		t,
		func(x *service_tester.Services) (any, any, error) {
			return x.Pattern.FindLiterals(o.Directory, "example/alfa", "Server")
		},
		o.Services,
		full,
	)
	assert.True(t, strings.Contains(result, "echo/pool.go"))
	assert.True(t, strings.Contains(result, "charlie/unit/only_test.go"))
}

func TestIndexedListCallsMatchesFullLoad(t *testing.T) {
	o, full := rewrite_tester.IndexedAndFull(t)
	rewrite_tester.SameResult(
		t,
		func(x *service_tester.Services) (any, any, error) {
			return x.Reference.ListCalls(o.Directory, "example/bravo", 20)
		},
		o.Services,
		full,
	)
}
