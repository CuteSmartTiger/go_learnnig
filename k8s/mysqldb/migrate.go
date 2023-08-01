package migrate

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateToDB(migrateDir string, dbSite string) error {
	db, _ := sql.Open("mysql", dbSite+"&multiStatements=true")
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://"+migrateDir,
		"mysql",
		driver,
	)

	err := m.Up()
	return err

}
