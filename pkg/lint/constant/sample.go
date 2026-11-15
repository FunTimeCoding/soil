package constant

const (
	FixtureExemptedSample  = "package example\n\n// golint:fixture stray_constant\nconst Foo = 1\n"
	FixtureMisplacedSample = "package example\n\n// golint:fixture stray_constant\n\nconst Foo = 1\n"
	FixtureDanglingSample  = "package example\n\n// golint:fixture stray_constant\n"
)
