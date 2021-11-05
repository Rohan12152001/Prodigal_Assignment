package load

import "github.com/Rohan12152001/Prodigal_assignment/data"

type LoadManager interface {
	LoadData(rowData []data.MfData) error
}