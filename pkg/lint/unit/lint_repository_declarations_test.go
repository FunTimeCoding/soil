package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"testing"
)

func TestRepositoryDeclaresPromotedMembers(t *testing.T) {
	r := shopRepository(t)
	assert.True(t, r.Declares("pkg/shop", "Item.Price"))
	assert.True(t, r.Declares("pkg/shop", "Item.Base"))
	assert.True(t, r.Declares("pkg/shop", "Item.Label"))
	assert.True(t, r.Declares("pkg/shop", "Item.Name"))
	assert.False(t, r.Declares("pkg/shop", "Item.Discount"))
	assert.False(
		t,
		r.Declares(
			"pkg/shop",
			join.Empty("Item", stringsConstant.Dot, constant.AnyField),
		),
	)
	assert.True(
		t,
		r.Declares(
			"pkg/shop",
			join.Empty("Locked", stringsConstant.Dot, constant.AnyField),
		),
	)
}

func TestRepositoryDeclaresTypesApartFromValues(t *testing.T) {
	r := shopRepository(t)
	assert.True(t, r.DeclaresType("pkg/shop", "Item"))
	assert.True(t, r.Declares("pkg/shop", "Default"))
	assert.False(t, r.DeclaresType("pkg/shop", "Default"))
}

func TestRepositoryKnowsTestNames(t *testing.T) {
	r := shopRepository(t)
	assert.True(t, r.Tests("TestShopKeepsStock"))
	assert.False(t, r.Tests("TestShopLosesStock"))
}

func TestRepositoryKnowsWhereTestNameRepeats(t *testing.T) {
	r := shopRepository(t)
	assert.True(t, r.TestElsewhere("TestShopKeepsStock", "pkg/shop/unit"))
	assert.False(t, r.TestElsewhere("TestStockCounts", "pkg/stock/unit"))
}
