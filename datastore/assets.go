package datastore

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/matt-FFFFFF/bookdata-api/loader"
)

// reuse struct from bookstore
type Assets struct {
	Store *[]*loader.BookData `json:"store"`
}

func (a *Assets) InitializeA() [][]string {

	//open file
	csvfile, err := os.Open("./assets/books.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	csv := csv.NewReader(csvfile)
	records, err := csv.ReadAll() // all CSV files into records
	if err != nil {
		log.Fatalln("Couldn't read the csv file", err)
	}
	return records

	//put it in the struct / variable

}
