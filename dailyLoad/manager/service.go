package dailyLoadExtract

import (
	"fmt"
	"github.com/Rohan12152001/Prodigal_assignment/dailyLoad/manager/db"
	extract "github.com/Rohan12152001/Prodigal_assignment/transform"
	"github.com/Rohan12152001/Prodigal_assignment/utils"
	_ "github.com/Rohan12152001/Prodigal_assignment/utils"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

type manager struct {
	db db.DailyLoadExtractDBManager
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

	// Format date initially
	startDate, err := m.db.FetchDate()
	if err != nil && err!=utils.NoRowsFound{
		logrus.Error(err)
		return err
	}
	endDate := time.Now().AddDate(0,0,-1)

	startTime := time.Now().Unix()

	// Call api
	formatUrlString := fmt.Sprintf("http://portal.amfiindia.com/DownloadNAVHistoryReport_Po.aspx?frmdt=%s&todt=%s",
		startDate.Format("02-Jan-2006"),
		endDate.Format("02-Jan-2006"))

	for getCallNumber:=1;getCallNumber<=10;getCallNumber++ {
		// Will retry at-most 10 times
		response, err := http.Get(formatUrlString)
		if response.StatusCode!=200{
			if getCallNumber==10{
				logrus.Errorf("The HTTP request failed, number of attempts exceeded!")
				return nil
			}
			continue
		}
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
			// Pass to transformService
			err := m.transformManager.TransformData(string(data))
			if err != nil {
				logrus.Error(err, "Couldn't transform data! (extract service)")
				return err
			}
			logrus.Infof("Data inserted from date: %s to date: %s", startDate.Format("02-Jan-2006"), endDate.Format("02-Jan-2006"))
		}

		// Update date in TimeDB
		m.db.UpdateDate(time.Now())
		logrus.Infof("New date in timeDB: %s", time.Now().Format("02-01-2006"))
	}

	logrus.Info("Total Time taken: ", time.Now().Unix() - startTime)
	return nil
}