package extract

import (
	"fmt"
	"github.com/Rohan12152001/Prodigal_assignment/extract/db"
	extract "github.com/Rohan12152001/Prodigal_assignment/transform"
	"github.com/Rohan12152001/Prodigal_assignment/utils"
	_ "github.com/Rohan12152001/Prodigal_assignment/utils"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

type manager struct {
	db db.ExtractDBManager
	transformManager extract.TransformManager
}

func New() ExtractManager {
	return manager{
		db: db.New(),
		transformManager: extract.New(),
	}
}

var logger = logrus.New()

func (m manager) ExtractData() error {
	fmt.Println("Starting the extraction...")
	fmt.Println("Please wait...")

	previousDate, err := m.db.FetchDate()
	if err != nil && err!=utils.NoRowsFound{
		logrus.Error(err)
		return err
	}

	// If date is 1st april 2006 >> Initial Load
	startTime := time.Now().Unix()
	if err==utils.NoRowsFound{
		// Format date initially
		startDate := time.Date(2006, time.April, 01, 0,0,0,0,time.UTC)
		endDate := startDate.AddDate(0,0, 89)

		// Loop and call for all data
		for true{
			logrus.Infof("Reading from date: %s to date: %s", startDate.Format("02-Jan-2006"), endDate.Format("02-Jan-2006"))
			// Call api
			formatUrlString := fmt.Sprintf("http://portal.amfiindia.com/DownloadNAVHistoryReport_Po.aspx?frmdt=%s&todt=%s",
											startDate.Format("02-Jan-2006"),
											endDate.Format("02-Jan-2006"))
			response, err := http.Get(formatUrlString)
			if err != nil {
				logrus.Errorf("The HTTP request failed with error %s\n", err)
				return err
			} else {
				if response.Body==nil{
					continue
				}
				logrus.Info("Data received!")
				data, _ := ioutil.ReadAll(response.Body)
				response.Body.Close()
				// Pass ahead
				err := m.transformManager.TransformData(string(data))
				if err != nil {
					logrus.Error(err, "Couldn't transform data! (extract service)")
					return err
				}
				logrus.Infof("Data inserted from date: %s to date: %s", startDate.Format("02-Jan-2006"), endDate.Format("02-Jan-2006"))
			}

			// end loop when all inserted
			if endDate.Format("02-01-2006") == time.Now().AddDate(0,0,-1).Format("02-01-2006"){
				m.db.UpdateDate(time.Now().Format("02-01-2006"))
				break
			}
			startDate = endDate.AddDate(0,0,1)
			endDate = startDate.AddDate(0,0,89)
			if endDate.After(time.Now()){
				endDate = time.Now().AddDate(0,0,-1)
			}
		}
	}else{
		// Format date initially
		startDate := previousDate
		endDate := time.Now().AddDate(0,0,-1)

		// Call api
		formatUrlString := fmt.Sprintf("http://portal.amfiindia.com/DownloadNAVHistoryReport_Po.aspx?frmdt=%s&todt=%s",
			startDate.Format("02-Jan-2006"),
			endDate.Format("02-Jan-2006"))
		response, err := http.Get(formatUrlString)
		if err != nil {
			logrus.Errorf("The HTTP request failed with error %s\n", err)
			return err
		} else {
			if response.Body==nil{
				logrus.Infof("No data received for frmdt=%s&todt=%s",startDate.Format("02-Jan-2006"),endDate.Format("02-Jan-2006"))
				return nil
			}
			logrus.Info("Data received!")
			data, _ := ioutil.ReadAll(response.Body)
			// Pass to transformation 239975
			err := m.transformManager.TransformData(string(data))
			if err != nil {
				logrus.Error(err)
				return err
			}
		}

		// Update date in TimeDB
		m.db.UpdateDate(time.Now().Format("02-01-2006"))
	}
	fmt.Println("Time took: ", time.Now().Unix() - startTime)

	return nil
}