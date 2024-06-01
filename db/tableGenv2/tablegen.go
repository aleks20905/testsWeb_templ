package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"slices"
	"sort"
	"strings"

	"github.com/aleks20905/testWeb_templ/db/config"

	"github.com/lib/pq"
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

	// Create a new table to store dynamically created table names
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS dynamic_tables (table_name VARCHAR(100) PRIMARY KEY)")
	if err != nil {
		log.Fatal(err)
	}

	//ShowColums(tableName)

	// Simulate inserting data into the new table
	//insertData(tableName)

	deviceName := "device_1" // fmt.Sprintf("device %d", i)
	insertDataForDev(deviceName, GenRandomData(10))

}

// func insertData(tableName string) {
// 	for i := 0; i < 5; i++ {
// 		insertRowQuery := fmt.Sprintf("INSERT INTO %s (send_at) VALUES ($1)", tableName)
// 		_, err := db.Exec(insertRowQuery, time.Now())
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println("Inserted new row into", tableName)
// 		time.Sleep(1 * time.Second) // Sleep for demonstration purpose
// 	}

// }
// func ShowColums(tableName string) {
// 	// Query and print metadata about columns in the new table
// 	var savedColumnNames []string
// 	err := db.QueryRow("SELECT column_names FROM table_columns WHERE table_name = $1", tableName).Scan(pq.Array(&savedColumnNames))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("Columns in table %s:\n", tableName)
// 	for _, columnName := range savedColumnNames {
// 		fmt.Println(columnName)
// 	}

// }
func simAddingColumn(tableName string) {
	// Simulate adding columns dynamically to the new table and store metadata about the columns
	columnNames := []string{"column_1", "column_2", "column_3"} //! aways updates the array / it doest adds or removes it just makes new array

	// Check if the table already exists in the table_columns table
	var existingColumns []string
	err := db.QueryRow("SELECT column_names FROM table_columns WHERE table_name = $1", tableName).Scan(pq.Array(&existingColumns))
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	// Convert existingColumns array to a map for easier comparison
	existingColumnsMap := make(map[string]struct{})
	for _, column := range existingColumns {
		existingColumnsMap[column] = struct{}{}
	}
	fmt.Println(existingColumnsMap)

	// Check if the existing and new column arrays are different
	var columnsChanged bool
	if len(existingColumns) != len(columnNames) {
		columnsChanged = true
	} else {
		for _, column := range columnNames {
			if _, exists := existingColumnsMap[column]; !exists {
				columnsChanged = true
				break
			}
		}
	}

	if columnsChanged {
		// Update the column array
		updateMetadataQuery := "UPDATE table_columns SET column_names = $1 WHERE table_name = $2"
		_, err = db.Exec(updateMetadataQuery, pq.Array(columnNames), tableName)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Updated metadata for table %s\n", tableName)
	} else {
		fmt.Printf("Metadata for table %s is up to date\n", tableName)
	}
}

/*
GENERATE random data for the 1 device.

Takes 1 argument how many sensors to generate
*/
func GenRandomData(n int) map[string]float32 {

	data := make(map[string]float32)

	//GEN RANDOM DATA FOR THE DEVICE  map[device 1:[{sensor 0 38.788177} {sensor 1 29.33254} {sensor 2 48.365208} {sensor 3 54.31514} {sensor 4 28.623657} {sensor 5 54.149876} {sensor 6 6.9542475} {sensor 7 35.740017} {sensor 8 43.792492} {sensor 9 24.22456}]]
	for i := range n {
		sensor := fmt.Sprintf("sensor_%d", i)
		data[sensor] = rand.Float32()*50 + 5
	}
	//fmt.Println(data) // show the data
	//GEN RANDOM DATA FOR THE DEVICE  map[device 1:[{sensor 0 38.788177} {sensor 1 29.33254} {sensor 2 48.365208} {sensor 3 54.31514} {sensor 4 28.623657} {sensor 5 54.149876} {sensor 6 6.9542475} {sensor 7 35.740017} {sensor 8 43.792492} {sensor 9 24.22456}]]
	return data
}

func insertDataForDev(deviceName string, data map[string]float32) {

	// CHECK IF deviceName EXIST
	err := CheckDeviceExist(deviceName)
	if err != nil {
		log.Fatal(err)
	}

	// CHECK IF all the provided Sensors EXIST
	err = CheckSensorExist(deviceName, data)
	if err != nil {
		log.Fatal(err)
	}

	//todo inser data
	err = inserIntoDb(deviceName, data)
	if err != nil {
		log.Fatal(err)
	}

}

// getDeviceSensors retrieves sensor names for a device from the database.
func GetDeviceSenosors(deviceName string) ([]string, error) {

	if db == nil {
		return nil, errors.New("database connection is nil")
	}

	// Prepare the SQL query
	query := fmt.Sprintf(`
        SELECT column_name
        FROM information_schema.columns
        WHERE table_name = '%s';
    `, deviceName)

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Check if rows is nil
	if rows == nil {
		return nil, errors.New("query result set is nil")
	}

	// Iterate through the result set and collect column names
	var columns []string
	for rows.Next() {
		var columnName string
		if err := rows.Scan(&columnName); err != nil {
			return nil, err
		}
		columns = append(columns, columnName)
	}

	// Check for errors during rows iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return columns, nil
}

// checkDeviceExist checks if a device exists in the database and creates it if not.
func CheckDeviceExist(deviceName string) error {

	if db == nil {
		return errors.New("database connection is nil")
	}

	// *############ CHECK IN dynamic_tables FOR THE SPECIFIC DEVICE ############
	// Check if the table exists in the database
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = 'public' AND table_name = $1", deviceName).Scan(&count)
	if err != nil {
		return err
	}
	// *############ CHECK IN dynamic_tables FOR THE SPECIFIC DEVICE ############

	// *############  CREATE THE NEW_TABLE
	// *############  Insert the name of the created table into the dynamic_tables table only if it doesn't already exist
	if count == 0 {

		// CREATE THE NEW_TABLE
		createTableQuery := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (send_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL PRIMARY KEY)", deviceName)
		_, err = db.Exec(createTableQuery)
		if err != nil {
			return err
		}
	}
	// *############ Insert the name of the created table into the dynamic_tables table only if it doesn't already exist

	return nil
}

// checkSensorExist checks if sensors for a device exist in the database and creates them if not. need to be use afeter CheckDeviceExist becase doest check if the table exist
func CheckSensorExist(deviceName string, data map[string]float32) error {

	//GET THE DEVICE SENSORS IN A ARRAY [send_at sensor_1 sensor_2 sensor_3 sensor_4]
	columns, err := GetDeviceSenosors(deviceName)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	// //fmt.Println("Columns for table", deviceName, ":", columns)
	//GET THE DEVICE SENSORS IN A ARRAY [send_at sensor_1 sensor_2 sensor_3 sensor_4]

	//CHECK EVERY IS EVERY SENSOR IN THE TABLE
	// Collect the sensor names that need to be created
	var sensorsToCreate []string
	for sensorName := range data {
		if !slices.Contains(columns, sensorName) {
			sensorsToCreate = append(sensorsToCreate, sensorName)
		}
	}
	sort.Strings(sensorsToCreate) // ! need fix this works fine but coude takes alot from the host
	// without [sensor_8 sensor_0 sensor_1 sensor_2 sensor_4 sensor_7 sensor_3 sensor_5 sensor_6 sensor_9]
	// with it [sensor_0 sensor_1 sensor_2 sensor_3 sensor_4 sensor_5 sensor_6 sensor_7 sensor_8 sensor_9]

	// If there are sensors to create, call CreateNewSensorColumns
	if len(sensorsToCreate) > 0 {
		err := CreateNewSensorColumns(deviceName, sensorsToCreate)
		if err != nil {
			return err
		}
		fmt.Println("Created new columns for sensors:", sensorsToCreate)
	}
	return nil
}

/*
Function to create new columns in the database table for sensors

	sensorsToCreate = {"sensor_2", "sensor_3", "sensor_4"}
	err := CreateNewSensorColumns(deviceName, sensorsToCreate)
	if err != nil {
		return err
	}
*/
func CreateNewSensorColumns(deviceName string, sensorNames []string) error {
	// Construct the SQL query to create new columns
	var createQuery string
	for _, sensorName := range sensorNames {
		createQuery += fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s NUMERIC(6,2);\n", deviceName, sensorName)
	}

	// Execute the SQL query to create new columns
	_, err := db.Exec(createQuery)
	if err != nil {
		return err
	}

	return nil
}

func inserIntoDb(deviceName string, data map[string]float32) error {

	if db == nil {
		return errors.New("database connection is nil")
	}

	// Constructing the INSERT statement with placeholders for each column
	var columns []string
	var placeholders []string
	var values []interface{}
	i := 1
	for sensorName, sensorValue := range data {
		columns = append(columns, sensorName)
		placeholders = append(placeholders, fmt.Sprintf("$%d", i))
		values = append(values, sensorValue)
		i++
	}
	stmt := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", deviceName, strings.Join(columns, ", "), strings.Join(placeholders, ", "))

	// Executing INSERT statement for each row
	_, err := db.Exec(stmt, values...)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Data inserted successfully to:", deviceName)

	return nil
}
