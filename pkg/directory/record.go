package directory

type Record struct {
	DistinguishedName string
	Attributes        map[string][]string
}
