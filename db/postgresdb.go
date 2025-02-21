package db

import (
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var log = logrus.New()

const MAX_DEFAULT_CONNECTION = 7

func ConnectDB() (*gorm.DB, error) {
	var err error
	connString := getConnectionString()

	// Open connection using GORM
	DB, err = gorm.Open(postgres.Open(connString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	sqlDB, err := DB.DB() // Get the underlying *sql.DB for connection settings
	if err != nil {
		log.Fatal("Failed to get sql.DB:", err)
	}

	// Set Connection Pooling
	dbConns := getDBConnection()
	sqlDB.SetMaxOpenConns(dbConns)             // Maximum open connections
	sqlDB.SetMaxIdleConns(dbConns)             // Maximum idle connections
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)  // Idle timeout
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // Max connection lifetime
	log.Info("Database connected successfully! with connection pooling")
	return DB, nil

}

func getConnectionString() string {
	// return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",os.Getenv("mydb_pg_user"),os.Getenv("mydb_pg_pw"),os.Getenv("mydb_pg_host"),os.Getenv("mydb_pg_dbname"))
	return "postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable"
}

func getDBConnection() int {
	dbConns := MAX_DEFAULT_CONNECTION
	if os.Getenv("MAX_DB_CONNS") != "" {
		env, err := strconv.Atoi(os.Getenv("MAX_DB_CONNS"))
		if err != nil {
			dbConns = env
		}
	}
	return dbConns
}
