package extract

import (
	"database/sql"
	"fmt"
	data2 "github.com/Rohan12152001/Prodigal_assignment/data"
	"github.com/Rohan12152001/Prodigal_assignment/load"
	"github.com/Rohan12152001/Prodigal_assignment/utils"
	_ "github.com/Rohan12152001/Prodigal_assignment/utils"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

type manager struct {
	loadManager load.LoadManager
}

func New() TransformManager {
	return manager{
		loadManager: load.New(),
	}
}

var logger = logrus.New()

func (m manager) TransformData(data string) error {
	parentMf := ""
	rows := []data2.MfData{}

	// parse
	actualRowCount := 0
	for _,line := range strings.Split(data, "\n"){
		splitArray := strings.Split(line, ";")

		// Check first element an int ?
		if schemeCode, err := strconv.Atoi(splitArray[0]); err == nil {
			actualRowCount += 1
			var nullableNav, nullableRepurchasePrice, nullableSalePrice sql.NullFloat64
			var nullableSchemeName, nullableIsinDivPayoutGrowth, nullableIsinDivReinvestment sql.NullString
			if nav, err := strconv.ParseFloat(splitArray[4], 64);err == nil{
				nullableNav.Scan(nav)
			}
			if repurchasePrice, err := strconv.ParseFloat(splitArray[4], 64);err == nil{
				nullableNav.Scan(repurchasePrice)
			}
			if salePrice, err := strconv.ParseFloat(splitArray[4], 64);err == nil{
				nullableNav.Scan(salePrice)
			}
			if splitArray[1]!=""{
				nullableSchemeName.Scan(splitArray[1])
			}
			if splitArray[2]!=""{
				nullableSchemeName.Scan(splitArray[2])
			}
			if splitArray[3]!=""{
				nullableSchemeName.Scan(splitArray[3])
			}

			date, err := time.Parse("02-Jan-2006", strings.TrimSuffix(splitArray[7], "\r"))
			if err != nil {
				fmt.Println(strings.TrimSuffix(splitArray[7], "\r"))
				logrus.Error(err, "Couldn't parse date! (transform service)")
				continue
			}

			// This is a row data
			rows = append(rows, data2.MfData{
				SchemeCode : schemeCode,
				SchemeName : nullableSchemeName,
				ParentScheme : parentMf,
				IsinDivPayoutGrowth : nullableIsinDivPayoutGrowth,
				IsinDivReinvestment : nullableIsinDivReinvestment,
				Nav : nullableNav,
				RepurchasePrice : nullableRepurchasePrice,
				SalePrice : nullableSalePrice,
				Date : date,
			})
		}
		// Change MF if Parent name present?
		if a := strings.Split(splitArray[0], " "); utils.StringInSlice("Mutual", a){
			parentMf = splitArray[0]
		}
	}
	// Also, insert if rows formed!
	if parentMf!=""{
		logrus.Info("Rows formed, calling Load data!")
		// Insert call
		err := m.loadManager.LoadData(rows)
		if err != nil {
			logrus.Error(err, "Couldn't Load data (transform service)")
			return err
		}
	}

	logrus.Infof("Actual rows count: %d", actualRowCount)
	return nil
}