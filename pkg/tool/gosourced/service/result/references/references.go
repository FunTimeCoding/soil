package references

import "github.com/funtimecoding/soil/pkg/tool/gosourced/service/result/location"

type References struct {
	Symbol    string               `json:"symbol"`
	Total     int                  `json:"total"`
	More      int                  `json:"more"`
	Locations []*location.Location `json:"locations"`
}
