package constant

const FixtureModule = `module fixture

go 1.27
`

const FixtureShape = `package shape

const Ceiling = 100

func Grade(score int) string {
	if score < 0 {
		return "invalid"
	}

	if score >= 90 {
		return "high"
	}

	return "low"
}

func Label(score int) string {
	return Grade(score) + "!"
}

func Ratio(a int, b int) int {
	return a / b
}
`

const FixtureShapeTest = `package shape

import "testing"

func TestGradeHigh(t *testing.T) {
	if Grade(95) != "high" {
		t.Errorf("want high")
	}
}

func TestGradeLow(t *testing.T) {
	if Grade(10) != "low" {
		t.Errorf("want low")
	}
}

func TestLabelRuns(t *testing.T) {
	_ = Label(50)
}

func TestRatio(t *testing.T) {
	if Ratio(6, 3) != 2 {
		t.Errorf("want 2")
	}
}
`

const FixtureBreaks = `packages: [./shape]

breaks:
  - name: high_bound
    expect: red
    edits:
      - file: shape/shape.go
        find: score >= 90
        replace: score > 900

  - name: zero_divisor
    expect: red
    edits:
      - file: shape/shape.go
        find: a / b
        replace: a / (b - b)

  - name: invalid_word
    expect: survive
    edits:
      - file: shape/shape.go
        find: return "invalid"
        replace: return "broken"

  - name: label_suffix
    expect: survive
    edits:
      - file: shape/shape.go
        find: '+ "!"'
        replace: '+ "?"'

  - name: absent_text
    expect: red
    edits:
      - file: shape/shape.go
        find: nothing like this
        replace: anything

  - name: unused_local
    expect: red
    edits:
      - file: shape/shape.go
        find: 'if score < 0 {'
        replace: 'if unused := 1; score < 0 {'

  - name: init_panic
    expect: red
    edits:
      - file: shape/shape.go
        find: 'const Ceiling = 100'
        replace: "const Ceiling = 100\n\nfunc init() { panic(\"setup\") }"

  - name: ceiling_value
    expect: survive
    edits:
      - file: shape/shape.go
        find: 'const Ceiling = 100'
        replace: 'const Ceiling = 200'

  - name: invalid_through_bound
    expect: survive
    edits:
      - file: shape/shape.go
        find: "return \"invalid\"\n\t}\n\n\tif score >= 90 {"
        replace: "return \"broken\"\n\t}\n\n\tif score >= 90 {"

  - name: empty_find
    expect: red
    edits:
      - file: shape/shape.go
        find: ''
        replace: anything
`

const FixtureNoTestsBreaks = `packages: [./shape]
tests: Nothing
`

const FixtureFailingTest = `package shape

import "testing"

func TestAlreadyFails(t *testing.T) {
	t.Errorf("fails before any break")
}
`
