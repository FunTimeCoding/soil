package mock_client

type Client struct {
	Failure error
	Closed  int
}
