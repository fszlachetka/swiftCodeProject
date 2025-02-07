package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// GET /v1/swift-codes/:swift-code
func GetSwiftCode(c *gin.Context, db *sql.DB) {
	code := strings.ToUpper(c.Param("swift-code"))
	fmt.Println(code)

	var swiftEntry SwiftCode
	query := `SELECT countryISO2, swift_code, bankName, address, country_name, is_headquarter FROM swift_codes WHERE swift_code = $1`

	err := db.QueryRow(query, code).Scan(
		&swiftEntry.CountryISO2,
		&swiftEntry.SwiftCode,
		&swiftEntry.BankName,
		&swiftEntry.Address,
		&swiftEntry.CountryName,
		&swiftEntry.IsHeadquarter,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "SWIFT code not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if swiftEntry.IsHeadquarter {
		var branches []SwiftCode
		branchQuery := `SELECT countryISO2, swift_code, bankName, address, country_name, is_headquarter FROM swift_codes WHERE swift_code LIKE $1 AND swift_code <> $2`
		rows, err := db.Query(branchQuery, code[:8]+"%", code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch branches"})
			return
		}
		defer rows.Close()

		for rows.Next() {
			var branch SwiftCode
			if err := rows.Scan(
				&branch.CountryISO2,
				&branch.SwiftCode,
				&branch.BankName,
				&branch.Address,
				&branch.CountryName,
				&branch.IsHeadquarter,
			); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan branch "})
				return
			}
			branches = append(branches, branch)
		}

		c.JSON(http.StatusOK, gin.H{
			"address":       swiftEntry.Address,
			"bankName":      swiftEntry.BankName,
			"countryISO2":   swiftEntry.CountryISO2,
			"countryName":   swiftEntry.CountryName,
			"isHeadquarter": swiftEntry.IsHeadquarter,
			"swiftCode":     swiftEntry.SwiftCode,
			"branches":      branches,
		})
	} else {
		c.JSON(http.StatusOK, swiftEntry)
	}
}

// GET /v1/swift-codes/country/:countryISO2
func GetSwiftCodesByCountry(c *gin.Context, db *sql.DB) {
	countryCode := strings.ToUpper(c.Param("countryISO2"))
	rows, err := db.Query(`SELECT countryISO2, swift_code, bankName, address, country_name, is_headquarter FROM swift_codes WHERE countryISO2 = $1`, countryCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query failed"})
		return
	}
	defer rows.Close()

	var swiftCodes []SwiftCode
	for rows.Next() {
		var swiftEntry SwiftCode
		if err := rows.Scan(
			&swiftEntry.CountryISO2,
			&swiftEntry.SwiftCode,
			&swiftEntry.BankName,
			&swiftEntry.Address,
			&swiftEntry.CountryName,
			&swiftEntry.IsHeadquarter,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan row"})
			return
		}
		swiftCodes = append(swiftCodes, swiftEntry)
	}

	if len(swiftCodes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No SWIFT codes found for this country"})
		return
	}

	c.JSON(http.StatusOK, swiftCodes)
}

// POST /v1/swift-codes
func CreateSwiftCode(c *gin.Context, db *sql.DB) {
	var record SwiftCode
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(record)
	query := `INSERT INTO swift_codes (countryISO2, swift_code, bankName, address, country_name, is_headquarter) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := db.Exec(query,
		strings.ToUpper(record.CountryISO2),
		strings.ToUpper(record.SwiftCode),
		record.BankName,
		record.Address,
		strings.ToUpper(record.CountryName),
		record.IsHeadquarter,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create record"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"swiftCode": record.SwiftCode,
		"message":   "SWIFT code added successfully",
	})
}

// DELETE /v1/swift-codes/:swift-code
func DeleteSwiftCode(c *gin.Context, db *sql.DB) {
	code := strings.ToUpper(c.Param("swift-code"))

	query := `DELETE FROM swift_codes WHERE swift_code = $1`
	result, err := db.Exec(query, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete record"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "SWIFT code not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "SWIFT code deleted successfully"})
}
