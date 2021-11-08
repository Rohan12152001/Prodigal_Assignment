package main

import (
	initialLoadExtract "github.com/Rohan12152001/Prodigal_assignment/initialLoad/manager"
)

func main(){
	// Extract
	extObject := initialLoadExtract.New()
	err := extObject.ExtractData()
	if err != nil {
		return
	}
}