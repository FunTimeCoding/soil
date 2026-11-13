package block

import "github.com/funtimecoding/soil/pkg/text/constant"

func NewSearch(
	number int,
	search string,
	replace string,
) *Block {
	return &Block{
		Number:  number,
		Opening: constant.SearchMarker,
		Search:  search,
		Replace: replace,
	}
}
