package main

import (
	"log"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"

	"github.com/aleks20905/testWeb_templ/db"
)

// SensorData represents the structure of sensor data
type SensorData struct {
	SensorName string
	Data       float64
}

// getData retrieves sensor data from the database
func getData() []SensorData {
	// Fetch data from the database using your db package
	data, err := db.GetDataByDevice("Device 2")
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	// Convert data to SensorData struct
	var sensorData []SensorData
	for _, d := range data {
		sensorData = append(sensorData, SensorData{
			SensorName: d.SensorName,
			Data:       d.Data,
		})
	}

	return sensorData
}

func main() {
	// Get sensor data
	sensorData := getData()

	// Create a map to store series data for each sensor
	seriesData := make(map[string][]opts.LineData)

	// Populate map with data from SensorData
	for _, data := range sensorData {
		seriesData[data.SensorName] = append(seriesData[data.SensorName], opts.LineData{Value: data.Data})
	}

	// Create a line chart
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Sensor Data Line Chart",
		}),
	)

	// Add series for each sensor
	for sensorName, data := range seriesData {
		line.AddSeries(sensorName, data)
	}

	// Create a web page to embed the chart
	page := components.NewPage()
	page.AddCharts(line)

	// Serve the web page
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		err := page.Render(w)
		if err != nil {
			log.Println(err)
		}
	})

	log.Println("Server started at :8080")
	// Start the HTTP server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
