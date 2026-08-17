package response

type Address struct {
	Value string `json:"ipaddr"`
	Bits  int    `json:"subnetbits"`
}
