package caller

import "example/pkg/target"

func current() string {
	return "v1"
}

func Value() func(string) string {
	return target.Valued
}

func Call() string {
	return target.Called(current())
}

func Local() string {
	version := current()

	return target.Local(version)
}
