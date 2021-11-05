package db

import "github.com/Rohan12152001/Prodigal_assignment/data"

type LoadDBManager interface {
	LoadDataDB (rowData []data.MfData) error
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "MutualFund_DB"
)