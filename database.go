package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func insertRecord(record swiftInfo, db *sql.DB) {
	// Przygotowanie zapytania SQL
	sqlStatement := `INSERT INTO swift_codes (countryISO2, swift_code, bankName, address, country_name, is_headquarter) 
					 VALUES ($1, $2, $3, $4, $5, $6) 
					 ON CONFLICT (swift_code) DO NOTHING;`

	// Wykonanie zapytania z wartościami
	_, err := db.Exec(sqlStatement, record.countryISO2, record.swiftCode, record.bankName, record.address, record.countryName, record.isHeadquarter)
	if err != nil {
		log.Fatal(err)
	}
}

func connectToDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}
