package ultil

import (
	"time"
)

// GetCurrentMySQLDate get current datetime with mysql format
func GetCurrentMySQLDate() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}
