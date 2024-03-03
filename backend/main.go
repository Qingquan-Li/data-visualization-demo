package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// Person represents a record from the CSV file
type Person struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	City        string `json:"city"`
	County      string `json:"county"`
	State       string `json:"state"`
	Zip         string `json:"zip"`
	Phone1      string `json:"phone1"`
	Phone2      string `json:"phone2"`
	Email       string `json:"email"`
	Web         string `json:"web"`
}

var people []Person

func main() {
	// Load data from CSV file
	loadData()

	// Create a Gin router with default middleware
	router := gin.Default()
	// Add CORS middleware
	router.Use(CORSMiddleware())

	// Define routes
	router.GET("/api/data", getData)
	router.POST("/api/search", searchData)
	router.GET("/api/num-of-people-per-state", getNumOfPeoplePerState)

	// Start the server
	// Run on port 8080 by default
	// router.Run(":8080")
	router.Run()
}

func loadData() {
	// Open the CSV file
	file, err := os.Open("data/us-500.csv")
	if err != nil {
		panic("Failed to open CSV file: " + err.Error())
	}
	// defer the closing of the file
	defer file.Close()

	// Create a new reader
	reader := csv.NewReader(file)
	reader.Comma = ','             // CSV fields are separated by commas
	reader.LazyQuotes = true       // Allow quotes in a field
	reader.TrimLeadingSpace = true // Trim leading space

	// Read all the records
	records, err := reader.ReadAll()
	if err != nil {
		panic("Failed to read CSV file: " + err.Error())
	}

	if len(records) <= 1 {
		panic("No data in CSV file")
	}

	// Print the number of records
	fmt.Println("Number of records found:", len(records)-1)

	for _, record := range records[1:] { // Skip the header row
		people = append(people, Person{
			FirstName:   record[0],
			LastName:    record[1],
			CompanyName: record[2],
			Address:     record[3],
			City:        record[4],
			County:      record[5],
			State:       record[6],
			Zip:         record[7],
			Phone1:      record[8],
			Phone2:      record[9],
			Email:       record[10],
			Web:         record[11],
		})
	}
}

func CORSMiddleware() gin.HandlerFunc {
	allowedOrigins := []string{
		// Vite dev server port: 5173
		"http://localhost:5173",
		"http://192.168.0.101:5173",
		// Kubernetes node port: 30011
		"http://localhost:30011",
		// Domain names
		"https://data-visualization-demo.qingquanli.com",
		"https://csv-data-visualization.qingquanli.com",
	}

	return func(c *gin.Context) {
		// Allow all origins
		// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// Allow specific origin
		origin := c.Request.Header.Get("Origin")
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}

		// Allow all methods
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		// Allow all headers
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func getData(c *gin.Context) {
	c.JSON(http.StatusOK, people)
}

func searchData(c *gin.Context) {
	var search struct {
		SearchField string `json:"search_field"`
		TargetValue string `json:"target_value"`
	}

	if err := c.ShouldBindJSON(&search); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Search for the target value in the specified field
	// var results []Person
	// Initialize as an empty slice to avoid returning null
	results := []Person{}
	for _, person := range people {
		value := getFieldValue(&person, search.SearchField)
		if strings.EqualFold(value, search.TargetValue) {
			results = append(results, person)
		}
	}

	c.JSON(http.StatusOK, results)
}

func getFieldValue(p *Person, field string) string {
	switch field {
	case "first_name":
		return p.FirstName
	case "last_name":
		return p.LastName
	case "company_name":
		return p.CompanyName
	case "address":
		return p.Address
	case "city":
		return p.City
	case "county":
		return p.County
	case "state":
		return p.State
	case "zip":
		return p.Zip
	case "phone1":
		return p.Phone1
	case "phone2":
		return p.Phone2
	case "email":
		return p.Email
	case "web":
		return p.Web
	default:
		return ""
	}
}

func getNumOfPeoplePerState(c *gin.Context) {
	stateCounts := make(map[string]int)
	for _, person := range people {
		stateCounts[person.State]++
	}
	c.JSON(http.StatusOK, stateCounts)
}
