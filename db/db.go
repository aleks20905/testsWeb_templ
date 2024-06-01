package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/aleks20905/testWeb_templ/db/config"
	"github.com/aleks20905/testWeb_templ/db/model"
	_ "github.com/lib/pq"
)

var db *sql.DB

// not in use
func Connet() {
	var err error

	db, err = sql.Open("postgres", config.ConnectionString("gotest"))
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}
	defer db.Close()
	fmt.Println("db Connected")

}

func GetDataByDevice(deviceName string) ([]model.SensorData, error) {
	var err error

	db, err = sql.Open("postgres", config.ConnectionString("gotest"))
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}
	defer db.Close()

	query := `
        SELECT id, device, sensor_name, data, time 
        FROM sensor_data 
        WHERE device = $1
    `

	rows, err := db.Query(query, deviceName)
	if err != nil {
		return nil, fmt.Errorf("error querying data / wrong deviceName: %v", err)
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
