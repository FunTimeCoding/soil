package keychain_store

func New(
	service string,
	account string,
) *Store {
	return &Store{service: service, account: account}
}
