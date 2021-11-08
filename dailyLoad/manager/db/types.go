package db

import "time"

type DailyLoadExtractDBManager interface {
	FetchDate() (time.Time,error)
	UpdateDate(newDate time.Time) error
	SetDate(newDate time.Time) error
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "MutualFund_DB"
)