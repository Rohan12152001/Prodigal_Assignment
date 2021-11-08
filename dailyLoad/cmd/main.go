package main

import (
	dailyLoadExtract "github.com/Rohan12152001/Prodigal_assignment/dailyLoad/manager"
)

func main(){
	// Extract
	extObject := dailyLoadExtract.New()
	err := extObject.ExtractData()
	if err != nil {
		return
	}
}