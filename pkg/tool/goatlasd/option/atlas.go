package option

import "time"

type Atlas struct {
	Address          string
	MetricAddress    string
	ServiceTokens    []string
	LitePath         string
	PostgresLocator  string
	Issuer           string
	ClientIdentifier string
	ClientSecret     string
	EncryptionSecret string
	PublicLocator    string
	Interval         time.Duration
	Retention        time.Duration
}
