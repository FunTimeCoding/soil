package constant

const (
	FixtureExemptedSample  = "package example\n\n// golint:fixture stray_constant\nconst Foo = 1\n"
	FixtureMisplacedSample = "package example\n\n// golint:fixture stray_constant\n\nconst Foo = 1\n"
	FixtureDanglingSample  = "package example\n\n// golint:fixture stray_constant\n"
)

const SuppressCallSample = `package example

func f() {
	run(
		"alfa",
	) // goanalyze:ignore forbidden_call
	run("bravo") // goanalyze:ignore forbidden_call
	outer(
		inner(), // goanalyze:ignore forbidden_call
		"charlie",
	)
	guarded(func() {
		run("delta") // goanalyze:ignore forbidden_call
	})
}
`

const SuppressLiteralSample = `package example

func f() {
	_ = &Client{
		// goanalyze:ignore restricted_literal
		Transport: &Transport{Dial: d}, // goanalyze:ignore restricted_literal
	}
	_ = &Client{
		Transport: &Transport{
			Dial: d,
		}, // goanalyze:ignore restricted_literal
	} // goanalyze:ignore restricted_literal
	_ = &Client{
		Transport: &Transport{Dial: d}, // goanalyze:ignore restricted_literal
	}
}
`
