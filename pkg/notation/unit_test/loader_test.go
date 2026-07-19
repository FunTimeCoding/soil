package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/notation/loader"
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
				constant.Directory,
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
