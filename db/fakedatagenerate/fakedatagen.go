package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/aleks20905/testWeb_templ/db/config"
	_ "github.com/lib/pq"
)

var db *sql.DB // Global variable for the database object

func main() {
	var err error

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Open a connection to the database
	db, err = sql.Open("postgres", config.ConnectionString("gotest"))
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}
	defer db.Close()

	// Handle interrupt signal (Ctrl+C)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Infinite loop
	for {
		select {
		case <-interrupt:
			fmt.Println("Received interrupt signal. Closing database connection and stopping.")
			if err := db.Close(); err != nil {
				log.Fatalf("Error closing database connection: %v\n", err)
			}
			os.Exit(0)
		default:
			err = insertData() // Insert data into the database
			if err != nil {
				log.Fatalf("Error inserting data: %v\n", err)
			}
			fmt.Println("Data insertion successful")

			// Sleep for 2 seconds before next iteration
			time.Sleep(2 * time.Second)
		}
	}
}

func insertData() error {
	// Insert some sample data into the table
	insertStatement := `
        INSERT INTO sensor_data (device, sensor_name, data, time)
        VALUES ($1, $2, $3, $4)
    `
	currentTime := time.Now()

	// Generate random number of devices and sensors
	numDevices := rand.Intn(10) + 1
	numSensors := rand.Intn(10) + 1

	for i := 1; i <= numDevices; i++ {
		device := fmt.Sprintf("Device %d", i)
		for j := 1; j <= numSensors; j++ {
			sensorName := fmt.Sprintf("Sensor %d", j)
			data := rand.Float64() * 100 // Generate random data
			_, err := db.Exec(insertStatement, device, sensorName, data, currentTime)
			if err != nil {
				return fmt.Errorf("error inserting data: %v", err)
			}
			fmt.Printf("Inserted data for %s - %s successfully\n", device, sensorName)
		}
	}

	return nil
}
