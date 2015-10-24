package postgres

import (
	"github.com/jinzhu/gorm"
	//Blank import needed to call init from pq
	_ "github.com/lib/pq"
)

const (
	databaseType = "postgres"
)

// Config - struct that has the database configuration attrs
type Config struct {
	User    string
	DbName  string
	SSLMode string
}

// NewConfig returns a new Config object
func NewConfig(user, dbName, sslMode string) *Config {
	return &Config{
		User:    user,
		DbName:  dbName,
		SSLMode: sslMode,
	}
}

func buildArgs(config *Config) string {
	return "user=" + config.User + " dbname=" + config.DbName + " sslmode=" + config.SSLMode
}

//GetDatabase will return a pointer to gorm.DB base on the config struct passed as argument
func GetDatabase(config *Config) *gorm.DB {
	db, err := gorm.Open(databaseType, buildArgs(config))
	if err != nil {
		panic(err)
	}
	return &db
}
