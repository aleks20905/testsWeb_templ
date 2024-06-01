package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/aleks20905/testWeb_templ/db/config"
	"github.com/aleks20905/testWeb_templ/db/model"

	_ "github.com/lib/pq"
)

var db *sql.DB // Global variable for the database object

func main() {
	var err error

	// Open a connection to the database
	db, err = sql.Open("postgres", config.ConnectionString("gotest"))
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}
	defer db.Close()

	err = initializeDatabase() // Initialize DB
	if err != nil {
		log.Fatalf("Error initializing database: %v\n", err)
	}
	fmt.Println("Database initialization successful")

	err = insertData() // Insert data into the database
	if err != nil {
		log.Fatalf("Error inserting data: %v\n", err)
	}
	fmt.Println("Data insertion successful")

	deviceName := "Device 1" // Select device to get data from
	data, err := getDataByDevice(deviceName)
	if err != nil {
		log.Fatalf("Error getting data for %s: %v\n", deviceName, err)
	}

	// Print retrieved data
	for _, d := range data {
		fmt.Printf("ID: %d, Device: %s, Sensor Name: %s, Data: %.2f, Time: %s\n",
			d.ID, d.Device, d.SensorName, d.Data, d.Time.Format("2006-01-02 15:04:05")) // Adjust the timestamp format to match your template
	}

	// Example usage of sendDataToDB function
	err = sendDataToDB("Device 2", "Temperature", 30.0, time.Now())
	if err != nil {
		log.Fatalf("Error sending data to DB: %v\n", err)
	}
}

// initializeDatabase creates a new table if it doesn't exist
func initializeDatabase() error {
	// Create a new table
	createTableStatement := `
        CREATE TABLE IF NOT EXISTS sensor_data (
            id SERIAL PRIMARY KEY,
            device VARCHAR(50) NOT NULL,
            sensor_name VARCHAR(50) NOT NULL,
            data NUMERIC(6,2) NOT NULL,
            time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
        )
    `
	_, err := db.Exec(createTableStatement)
	if err != nil {
		return fmt.Errorf("error creating table: %v", err)
	}

	fmt.Println("Table created successfully")
	return nil
}

// insertData inserts sample data into the table
func insertData() error {
	// Insert some sample data into the table
	insertStatement := `
        INSERT INTO sensor_data (device, sensor_name, data, time)
        VALUES ($1, $2, $3, $4)
    `
	currentTime := time.Now()
	for i := 1; i <= 5; i++ {
		device := fmt.Sprintf("Device %d", i)
		sensorName := fmt.Sprintf("Sensor %d", i)
		data := float64(i) * 10
		_, err := db.Exec(insertStatement, device, sensorName, data, currentTime)
		if err != nil {
			return fmt.Errorf("error inserting data: %v", err)
		}
		fmt.Printf("Inserted data for %s successfully\n", device)
	}

	return nil
}

// getDataByDevice retrieves data for a specific device from the database
func getDataByDevice(deviceName string) ([]model.SensorData, error) {
	query := `
        SELECT id, device, sensor_name, data, time 
        FROM sensor_data 
        WHERE device = $1
    `

	rows, err := db.Query(query, deviceName)
	if err != nil {
		return nil, fmt.Errorf("error querying data: %v", err)
	}
	defer rows.Close()

	var results []model.SensorData
	for rows.Next() {
		var data model.SensorData
		err := rows.Scan(&data.ID, &data.Device, &data.SensorName, &data.Data, &data.Time)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		results = append(results, data)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return results, nil
}

// sendDataToDB inserts a new sensor data entry into the database
func sendDataToDB(device string, sensorName string, data float64, time time.Time) error {
	insertStatement := `
        INSERT INTO sensor_data (device, sensor_name, data, time)
        VALUES ($1, $2, $3, $4)
    `
	_, err := db.Exec(insertStatement, device, sensorName, data, time)
	if err != nil {
		return fmt.Errorf("error inserting data: %v", err)
	}
	fmt.Printf("Inserted data for Device: %s, Sensor Name: %s successfully\n", device, sensorName)
	return nil
}
