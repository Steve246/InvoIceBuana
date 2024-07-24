package manager

import (
	"fmt"
	"invoiceBuana/config"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Infra interface {
	SqlDb() *gorm.DB
}

type infra struct {
	dbResource *gorm.DB
}

func (i *infra) SqlDb() *gorm.DB {
	return i.dbResource
}

func NewInfra(config config.Config) Infra {

	resource, err := initDbResource(config.DataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print("Database Connected!")

	return &infra{dbResource: resource}

}

func initDbResource(dataSourceName string) (*gorm.DB, error) {

	// Open MySQL connection with GORM
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	fmt.Println("ini db --> ", db)

	if err != nil {
		return nil, err
	}

	env := os.Getenv("ENV")
	dbReturn := db
	if env == "migration" {
		dbReturn = db.Debug()

		// run migration pake go-migrate here

		migrationDirection := os.Getenv("MIGRATION_DIRECTION")

		err := runMigrations(dataSourceName, migrationDirection)
		if err != nil {
			return nil, err
		}

		//masukin table untuk dimigrate
	} else if env == "dev" {
		dbReturn = db.Debug()
	}
	if err != nil {
		return nil, err
	}
	return dbReturn, nil
}

func runMigrations(dataSourceName string, direction string) error {
	// Adjust DSN for golang-migrate
	migrateDSN := fmt.Sprintf("mysql://%s", dataSourceName)

	m, err := migrate.New(
		"file://db/migration",
		migrateDSN,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	if direction == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			return fmt.Errorf("failed to run migrations up: %w", err)
		}
	} else if direction == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			return fmt.Errorf("failed to run migrations down: %w", err)
		}
	} else {
		return fmt.Errorf("invalid migration direction: %s", direction)
	}

	return nil
}
