package load

import (
	"github.com/Rohan12152001/Prodigal_assignment/data"
	"github.com/Rohan12152001/Prodigal_assignment/load/db"
	_ "github.com/Rohan12152001/Prodigal_assignment/utils"
	"github.com/sirupsen/logrus"
)

type manager struct {
	db db.LoadDBManager
}

func New() LoadManager {
	return manager{
		db: db.New(),
	}
}

var logger = logrus.New()

func (m manager) LoadData(rowData []data.MfData) error {
	err := m.db.LoadDataDB(rowData)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

