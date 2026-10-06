package caller

import (
	"example/pkg/assert"
	"example/pkg/target"
)

func standalone() *target.Shape {
	a := &target.Shape{Draw: 1, Inhale: 2}

	return a
}

func again() *target.Shape {
	var b = &target.Shape{Inhale: 2, Draw: 1}

	return b
}

func superset() *target.Shape {
	c := &target.Shape{Draw: 1, Inhale: 2, Hold: 3}

	return c
}

func nested() []*target.Shape {
	return []*target.Shape{{Draw: 1}, &target.Shape{Hold: 4}}
}

func value() target.Shape {
	v := target.Shape{}

	return v
}

func allocated() *target.Shape {
	return new(target.Shape)
}

func unkeyed() *target.Shape {
	u := &target.Shape{1, 2, 3, 4}

	return u
}

func expected() bool {
	return assert.Equal(&target.Shape{Draw: 1}, standalone())
}
