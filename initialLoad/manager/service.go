package initialLoadExtract

import (
	"fmt"
	"github.com/Rohan12152001/Prodigal_assignment/initialLoad/manager/db"
	extract "github.com/Rohan12152001/Prodigal_assignment/transform"
	"github.com/Rohan12152001/Prodigal_assignment/utils"
	_ "github.com/Rohan12152001/Prodigal_assignment/utils"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

type manager struct {
	db db.InitialLoadExtractDBManager
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

	startTime := time.Now().Unix()

	var startDate, endDate time.Time

	// If no date found then start from April 2006
	if err==utils.NoRowsFound {
		// Format date initially
		startDate = time.Date(2006, time.April, 01, 0, 0, 0, 0, time.UTC)
		endDate = startDate.AddDate(0, 0, 89)
	}else{
		startDate = previousDate
		endDate = startDate.AddDate(0,0, 89)
		if endDate.After(time.Now()){
			endDate = time.Now().AddDate(0,0,-1)
		}
	}

	for true{
		logrus.Infof("Reading from date: %s to date: %s", startDate.Format("02-Jan-2006"), endDate.Format("02-Jan-2006"))
		// Call api
		formatUrlString := fmt.Sprintf("http://portal.amfiindia.com/DownloadNAVHistoryReport_Po.aspx?frmdt=%s&todt=%s",
			startDate.Format("02-Jan-2006"),
			endDate.Format("02-Jan-2006"))
		for getCallNumber:=1;getCallNumber<=10;getCallNumber++{
			// Will retry at-most 10 times
			logrus.Infof("Call number: %d", getCallNumber)
			response, err := http.Get(formatUrlString)
			if err != nil {
				if getCallNumber==10{
					logrus.Errorf("The HTTP request failed, number of attempts exceeded!")
					return nil
				}
				logrus.Errorf("The HTTP request failed with error %s\n", err)
				continue
			} else {
				if response.StatusCode!=200 || response.Body==nil{
					if getCallNumber==10{
						logrus.Errorf("The HTTP request failed, number of attempts exceeded!")
						return nil
					}
					continue
				}
				logrus.Info("Data received!")
				data, _ := ioutil.ReadAll(response.Body)
				response.Body.Close()
				// Pass ahead
				err := m.transformManager.TransformData(string(data))
				if err != nil {
					if err==utils.NoRowsFormed{
						if getCallNumber==10{
							logrus.Errorf("The HTTP request failed, number of attempts exceeded!")
							return nil
						}
						logrus.Info(err, "(extract service)")
						continue
					}
					logrus.Error(err, "Couldn't transform data! (extract service)")
					return err
				}
				logrus.Infof("Data inserted from date: %s to date: %s", startDate.Format("02-Jan-2006"), endDate.Format("02-Jan-2006"))
			}

			// Update date after every three months insert!
			if startDate.Format("02-01-2006") == "01-04-2006"{
				err := m.db.SetDate(endDate.AddDate(0,0,1))
				if err != nil {
					logrus.Error(err, " (extract service)")
					return err
				}
			}else{
				err := m.db.UpdateDate(endDate.AddDate(0,0,1))
				if err != nil {
					logrus.Error(err, " (extract service)")
					return err
				}
			}

			startDate = endDate.AddDate(0,0,1)
			endDate = startDate.AddDate(0,0,89)
			if endDate.After(time.Now()){
				endDate = time.Now().AddDate(0,0,-1)
			}
			break
		}

		// end loop when all inserted
		if endDate.Format("02-01-2006") == time.Now().AddDate(0,0,-1).Format("02-01-2006"){
			m.db.SetDate(time.Now())
			logrus.Infof("New date in timeDB: %s", time.Now().Format("02-01-2006"))
			break
		}
	}

	logrus.Info("Total Time taken: ", time.Now().Unix() - startTime)
	return nil
}