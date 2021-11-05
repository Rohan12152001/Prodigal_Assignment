package db

import (
	"fmt"
	"github.com/Rohan12152001/Prodigal_assignment/data"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type manager struct {
	db *sqlx.DB
}

func New() LoadDBManager {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return manager{
		db: db,
	}
}

// ReplaceSQL replaces the instance occurrence of any string pattern with an increasing $n based sequence
func ReplaceSQL(old, searchPattern string) string {
	tmpCount := strings.Count(old, searchPattern)
	for m := 1; m <= tmpCount; m++ {
		old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	}
	return old
}

func (m manager) LoadDataDB(rowData []data.MfData) error {
	for _, row := range rowData {
		sqlStr := "INSERT INTO MF_data (schemeCode, schemeName, parentScheme, isinDivPayoutGrowth, isinDivReinvestment, nav, repurchasePrice, salePrice, date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
		_, err := m.db.Exec(sqlStr,
				row.SchemeCode,
				row.SchemeName,
				row.ParentScheme,
				row.IsinDivPayoutGrowth,
				row.IsinDivReinvestment,
				row.Nav,
				row.RepurchasePrice,
				row.SalePrice,
				row.Date)
		if err != nil {
			return err
		}
	}
	logrus.Info("Inserted Data!")


	// TODO:  For batch insert OPTIMISE?
	//sqlStr := "INSERT INTO MF_data (schemeCode, schemeName, parentScheme, isinDivPayoutGrowth, isinDivReinvestment, nav, repurchasePrice, salePrice, date) VALUES "
	//vals := []interface{}{}
	//
	//for _, row := range rowData {
	//	//sqlStr += "($1,$2,$3,$4,$5,$6,$7,$8),"
	//	sqlStr += "(?,?,?,?,?,?,?,?,?),"
	//	vals = append(vals,
	//		row.SchemeCode,
	//		row.SchemeName,
	//		row.ParentScheme,
	//		row.IsinDivPayoutGrowth,
	//		row.IsinDivReinvestment,
	//		row.Nav,
	//		row.RepurchasePrice,
	//		row.SalePrice,
	//		row.Date)
	//}
	//
	////trim the last ,
	//sqlStr = strings.TrimSuffix(sqlStr, ",")
	//
	////Replacing ? with $n for postgres
	////sqlStr = ReplaceSQL(sqlStr, "?")
	//
	////prepare the statement
	//stmt, _ := m.db.Prepare(sqlStr)
	//
	////format all vals at once
	//logrus.Info("Calling Exec() in DB service!")
	//_, err := stmt.Exec(vals...)
	//if err != nil {
	//	return err
	//}
	//
	//stmt.Close()
	//logrus.Info("Inserted Data!")

	return nil
}