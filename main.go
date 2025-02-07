package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db, err := connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	parseDataFromExceltoDB("Interns_2025_SWIFT_CODES.xlsx", db)

	ginServer := gin.Default()
	ginGroup := ginServer.Group("/v1")
	{
		ginGroup.GET("/swift-codes/:swift-code", func(c *gin.Context) {
			GetSwiftCode(c, db)
		})
		ginGroup.GET("/swift-codes/country/:countryISO2", func(c *gin.Context) {
			GetSwiftCodesByCountry(c, db)
		})
		ginGroup.POST("/swift-codes", func(c *gin.Context) {
			CreateSwiftCode(c, db)
		})
		ginGroup.DELETE("/swift-codes/:swift-code", func(c *gin.Context) {
			DeleteSwiftCode(c, db)
		})
	}

	if err := ginServer.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}

}
