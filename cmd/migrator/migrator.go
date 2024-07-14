package migrator

import (
	"database/sql"
	"fmt"
	"os"
)

func Migrate(db *sql.DB, filePath string) error {
	migrationFile, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("cannot read migration file: %v", err)
	}

	_, err = db.Exec(string(migrationFile))
	if err != nil {
		fmt.Print(err)
		return fmt.Errorf("unable to apply migration")
	}

	return nil
}
