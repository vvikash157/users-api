package db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var log=config.InitializeLogger()

/*
Migrations applies all pending migrations
use this command to generate sql file
`migrate create -ext sql -dir ./.db/migrations add_new_columns_to_users`
*/
func Migrations() {
	//return if migration not needed
	if strings.ToLower(os.Getenv("IS_MIGRATIONS_NEED")) == "false" {
		return
	}

	db, err := sql.Open("postgres", getConnectionString())
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("Failed to create migration driver:", err)
	}

	migrationPath := "file://.db/migrations"

	m, err := migrate.NewWithDatabaseInstance(migrationPath, "postgres", driver)
	if err != nil {
		log.Fatal("Failed to initialize migration:", err)
	}

	version, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		log.Fatal("Unable to get migration version:", err)
	} else if err == migrate.ErrNilVersion {
		fmt.Println("No migrations applied yet.")
	} else {
		fmt.Printf("Current migration version: %d (dirty: %v)\n", version, dirty)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migration failed:", err)
	} else if err == migrate.ErrNoChange {
		fmt.Println("No new migrations to apply.")
	} else {
		fmt.Println("Database migrated successfully!")
	}
}

func GenerateVersionString(version uint, dirty bool, err error) string {
	if err != nil {
		return fmt.Sprintf("unable to get version info from go-migrate : %v", err)
	} else {
		return fmt.Sprintf("current version is %d and dirty flag is %v", version, dirty)
	}
}
