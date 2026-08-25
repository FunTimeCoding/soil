package mock_client

type Client struct {
	pages          map[string]*entry
	nextIdentifier int
}
