package search_index

import (
	"database/sql"
	"sync"
)

type Index struct {
	database *sql.DB
	mutex    sync.Mutex
}
