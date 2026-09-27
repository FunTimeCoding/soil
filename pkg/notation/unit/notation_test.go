package unit

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/notation/fixture"
	"github.com/funtimecoding/soil/pkg/notation/loader"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
	"testing"
)

func TestLoader(t *testing.T) {
	l := loader.New()
	l.Load(
		join.Absolute(
			system.FindDirectoryUp(
				system.WorkDirectory(),
				gitConstant.Directory,
			),
			systemConstant.FixturePath,
			systemConstant.NotationPath,
		),
	)
	assert.Count(t, 2, l.Contents())
	assert.Any(
		t,
		`{
    "Classified": "AnAlertName",
    "Reason": "A reason why the answer was chosen",
    "Answer": "not-yet-broken"
}
`,
		l.Contents()["response-1.json"],
	)
	actualMap := l.ToMap()
	assert.Any(
		t,
		map[string]string{
			"Classified": "AnAlertName",
			"Reason":     "A reason why the answer was chosen",
			"Answer":     "not-yet-broken",
		},
		actualMap["response-1.json"],
	)
	assert.Any(
		t,
		map[string]string{
			"Classified": "AnotherAlertName",
			"Reason":     "Another reason why the answer was chosen",
			"Answer":     "already-broken",
		},
		actualMap["response-2.json"],
	)
	reducedMap := actualMap
	delete(reducedMap, "response-2.json")
	assert.String(
		t,
		`1
Answer: not-yet-broken
Classified: AnAlertName
Reason: A reason why the answer was chosen`,
		l.ToText(reducedMap),
	)
}

func TestComment(t *testing.T) {
	assert.String(t, "a:\n\"b\"", notation.Comment("a", "b"))
}

func TestDecodeAny(t *testing.T) {
	var a any
	notation.DecodeAny(true, &a)
	assert.Any(t, a, true)
	var b any
	notation.DecodeAny(constant.UpperAlfa, &b)
	assert.Any(t, b, "Alfa")
}

func TestDecode(t *testing.T) {
	var actual []int
	assert.FatalOnError(t, notation.Decode("[1]", &actual))
	assert.Any(t, []int{1}, actual)
}

func TestDecodeStrict(t *testing.T) {
	var actual []int
	notation.MustDecode("[1]", &actual, false)
	assert.Any(t, []int{1}, actual)
}

func TestUnknown(t *testing.T) {
	raw := `{
		"name": "jdoe",
		"department": "Development",
		"location": "Earth",
		"skills": ["Go", "Kubernetes"]
	}`
	var u fixture.User
	errors.PanicOnError(json.Unmarshal([]byte(raw), &u))
	assert.String(t, "jdoe", u.Name)
	assert.Any(t, "Development", u.Unknown["department"])
	assert.Any(t, "Earth", u.Unknown["location"])
	assert.Any(t, []any{"Go", "Kubernetes"}, u.Unknown["skills"])
}

func TestFormatTypes(t *testing.T) {
	assert.Any(
		t,
		"{\n    \"String\": \"a\",\n    \"Integer\": 1,\n    \"Float\": 1.5,\n    \"Boolean\": true\n}",
		notation.Format(
			fixture.Primitives{
				String:  "a",
				Integer: 1,
				Float:   1.5,
				Boolean: true,
			},
		),
	)
}

func TestFormatStringWithVector(t *testing.T) {
	assert.Any(
		t,
		"{\n    \"String\": \"1,<1.0, 1.0, 1.0>,2\"\n}",
		notation.Format(fixture.WithString{String: "1,<1.0, 1.0, 1.0>,2"}),
	)
}

func TestMarshallIndent(t *testing.T) {
	assert.Any(t, `"a"`, notation.MarshalIndent("a"))
}

func TestWithoutRemovesKey(t *testing.T) {
	assert.String(
		t,
		"{\n\t\"Inner\": null,\n\t\"Name\": \"a\"\n}",
		notation.MarshalIndent(
			notation.Without(
				fixture.NewWrapped(
					"a",
					nil,
					fixture.NewPrimitives("b", 0, 0, false),
				),
				"Raw",
			),
		),
	)
}

func TestWithoutRemovesNestedKey(t *testing.T) {
	assert.String(
		t,
		"{\n\t\"Inner\": {\n\t\t\"Inner\": null,\n\t\t\"Name\": \"b\"\n\t},\n\t\"Name\": \"a\"\n}",
		notation.MarshalIndent(
			notation.Without(
				fixture.NewWrapped(
					"a",
					fixture.NewWrapped(
						"b",
						nil,
						fixture.NewPrimitives("d", 0, 0, false),
					),
					fixture.NewPrimitives("c", 0, 0, false),
				),
				"Raw",
			),
		),
	)
}

func TestWithoutRemovesKeyInSlice(t *testing.T) {
	assert.String(
		t,
		"[\n\t{\n\t\t\"Inner\": null,\n\t\t\"Name\": \"a\"\n\t},\n\t{\n\t\t\"Inner\": null,\n\t\t\"Name\": \"b\"\n\t}\n]",
		notation.MarshalIndent(
			notation.Without(
				[]*fixture.Wrapped{
					fixture.NewWrapped(
						"a",
						nil,
						fixture.NewPrimitives("c", 0, 0, false),
					),
					fixture.NewWrapped(
						"b",
						nil,
						fixture.NewPrimitives("d", 0, 0, false),
					),
				},
				"Raw",
			),
		),
	)
}

func TestWithoutKeepsUnmatchedKey(t *testing.T) {
	assert.String(
		t,
		"{\n\t\"Boolean\": false,\n\t\"Float\": 0,\n\t\"Integer\": 0,\n\t\"String\": \"a\"\n}",
		notation.MarshalIndent(
			notation.Without(fixture.NewPrimitives("a", 0, 0, false), "Raw"),
		),
	)
}

func TestWithoutScalar(t *testing.T) {
	assert.Any(t, "a", notation.Without("a", "Raw"))
}
