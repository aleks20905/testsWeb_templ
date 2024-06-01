package config

import "fmt"

const (
	Host     = "localhost"
	Port     = 5433
	User     = "postgres"
	Password = "postgres"
	DBName   = "gotest"
)

// ConnectionString returns the connection string for the PostgreSQL database.
// If a database name is provided, it uses that name; otherwise, it uses the default database name. That is set in db/config/config.go
//
// Example usage:
//   connStr := ConnectionString("mydatabase")
//   connStr := ConnectionString()
func ConnectionString(dbName ...string) string {
	if len(dbName) > 0 {
		return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			Host, Port, User, Password, dbName[0])
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, DBName)
}
