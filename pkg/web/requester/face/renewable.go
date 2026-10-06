package face

type Renewable interface {
	Authorizer
	Renew() error
}
