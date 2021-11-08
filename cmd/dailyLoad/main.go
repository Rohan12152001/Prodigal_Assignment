package main

import (
	"github.com/Rohan12152001/Prodigal_assignment/extract"
)

func main(){
	// Extract
	extObject := extract.New()
	err := extObject.ExtractData()
	if err != nil {
		return
	}

	// Cron
	//c := cron.New()
	//defer func() {
	//	c.Stop()
	//}()
	//////c.AddFunc("@daily", func() {
	//////	err = extObject.ExtractData()
	//////	if err != nil {
	//////		return
	//////	}
	//////})
	//c.AddFunc("5 * * * * *", func() {
	//	err = extObject.ExtractData()
	//	if err != nil {
	//		return
	//	}
	//})
	//c.Start()
	//time.Sleep(time.Second*100)
}