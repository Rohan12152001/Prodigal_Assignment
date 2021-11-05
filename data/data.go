package data

import (
	"database/sql"
	"time"
)

type MfData struct {
	SchemeCode int
	SchemeName sql.NullString
	ParentScheme string
	IsinDivPayoutGrowth sql.NullString
	IsinDivReinvestment sql.NullString
	Nav sql.NullFloat64
	RepurchasePrice sql.NullFloat64
	SalePrice sql.NullFloat64
	Date time.Time
}

