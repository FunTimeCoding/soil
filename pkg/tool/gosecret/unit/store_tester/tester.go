package store_tester

type Tester struct {
	stored   map[string]map[string]string
	existing map[string]string
	tamper   string
}
