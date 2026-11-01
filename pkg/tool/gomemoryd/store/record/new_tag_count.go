package record

func NewTagCount(
	tag string,
	count int,
) *TagCount {
	return &TagCount{Tag: tag, Count: count}
}
