package db

import (
	"database/sql"
	"fmt"
	"github.com/Rohan12152001/Prodigal_assignment/utils"
	"github.com/jackc/pgx/pgtype"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"time"
)

type manager struct {
	db *sqlx.DB
}

func New() InitialLoadExtractDBManager {
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

func (m manager) FetchDate() (time.Time,error) {
	query := "Select DATE (fetchFrom) as fetchFrom from timeStampDB"

	var dt pgtype.Date
	err := m.db.QueryRow(query).Scan(&dt)
	if err != nil {
		if err == sql.ErrNoRows{
			return time.Time{}, utils.NoRowsFound
		}
		return time.Time{}, err
	}

	return dt.Time, nil
}

func (m manager) UpdateDate(newDate time.Time) error {
	updateQuery := "UPDATE timeStampDB SET fetchFrom=$1;"

	_, err := m.db.Exec(updateQuery, newDate)
	if err != nil {
		return err
	}

	logrus.Info("Date updated!")

	return nil
}

func (m manager) SetDate(newDate time.Time) error {
	updateQuery := "INSERT INTO timeStampDB VALUES ($1);"

	_, err := m.db.Exec(updateQuery, newDate)
	if err != nil {
		return err
	}

	logrus.Info("Date inserted!")

	return nil
}