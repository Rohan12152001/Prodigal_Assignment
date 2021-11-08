package utils

import "fmt"

var (
	NoRowsFound = fmt.Errorf("no rows found")
	NoRowsFormed = fmt.Errorf("AFMI gave wrong response, retry")
)


