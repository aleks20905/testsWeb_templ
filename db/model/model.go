package model

import "time"

// The "Time" can be formated like:
//
// 		fmt.Printf("Time: %s\n", d.Time.Format("2006-01-02 15:04:05"))
//
type SensorData struct {
	ID         int
	Device     string
	SensorName string
	Data       float64
	Time       time.Time
}
