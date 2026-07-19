package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/structs/notation_tag"
	"github.com/funtimecoding/soil/pkg/structs/relational_tag"
	"reflect"
	"testing"
)

func TestPrimary(t *testing.T) {
	type Test struct {
		Name string `json:"name" gorm:"primaryKey"`
	}
	v := reflect.TypeOf(Test{})
	assert.Integer(t, 1, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		n := notation_tag.FromField(f)
		assert.String(t, "name", n.Key())
		assert.False(t, n.OmitEmpty())
		r := relational_tag.FromField(f)
		assert.True(t, r.Primary())
		assert.True(t, !r.Nullable())
	}
}

func TestOmit(t *testing.T) {
	type Test struct {
		Age int `json:"age,omitempty" gorm:"not null"`
	}
	v := reflect.TypeOf(Test{})
	assert.Integer(t, 1, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		n := notation_tag.FromField(f)
		assert.String(t, "age", n.Key())
		assert.True(t, n.OmitEmpty())
		r := relational_tag.FromField(f)
		assert.False(t, r.Primary())
		assert.True(t, !r.Nullable())
	}
}

func TestNotNullDefault(t *testing.T) {
	type Test struct {
		Mail string `json:"mail" gorm:"not null;default:jdoe@example.org"`
	}
	v := reflect.TypeOf(Test{})
	assert.Integer(t, 1, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		n := notation_tag.FromField(f)
		assert.String(t, "mail", n.Key())
		assert.False(t, n.OmitEmpty())
		r := relational_tag.FromField(f)
		assert.False(t, r.Primary())
		assert.True(t, !r.Nullable())
		assert.String(t, "jdoe@example.org", r.Default())
	}
}

func TestSkip(t *testing.T) {
	type Test struct {
		Unknown string `json:"-"`
	}
	v := reflect.TypeOf(Test{})
	assert.Integer(t, 1, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		assert.Nil(t, notation_tag.FromField(v.Field(i)))
	}
}
