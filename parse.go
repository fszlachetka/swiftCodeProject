package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/xuri/excelize/v2"
	"log"
)

func parseDataFromExceltoDB(fileName string, db *sql.DB) {
	file, err := excelize.OpenFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rows, err := file.GetRows(file.GetSheetName(0))
	if err != nil {
		log.Fatal(err)
	}

	for index, row := range rows {
		if index == 0 {
			continue
		}
		if len(row[0]) != 2 {
			continue
		}

		var record swiftInfo
		record.address = row[4]
		record.bankName = row[3]
		record.countryISO2 = row[0]
		record.countryName = row[5]
		record.swiftCode = row[1]
		code := record.swiftCode

		if code[len(code)-3:] == "XXX" {
			record.isHeadquarter = true
		} else {
			record.isHeadquarter = false
		}
		insertRecord(record, db)
	}
}
