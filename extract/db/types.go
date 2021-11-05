package db

import "time"

type ExtractDBManager interface {
	FetchDate() (time.Time,error)
	UpdateDate(newDate string) error
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "MutualFund_DB"
)