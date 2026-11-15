package result

func NewUsageMethod(
	name string,
	file string,
) *UsageMethod {
	return &UsageMethod{
		Name:   name,
		File:   file,
		Tables: NewUsageTables(nil, nil),
		Reach:  NewUsageTables(nil, nil),
	}
}
