package database

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)


func RunMigrations(
	ctx context.Context,
	db DB,
) error {


	_, err := db.ExecContext(
		ctx,
		`
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version TEXT PRIMARY KEY,
			applied_at TIMESTAMP DEFAULT NOW()
		)
		`,
	)

	if err != nil {
		return err
	}


	files, err := filepath.Glob(
		"migrations/*.sql",
	)

	if err != nil {
		return err
	}


	sort.Strings(files)


	for _, file := range files {


		version := filepath.Base(file)


		var exists bool


		err := db.QueryRowContext(
			ctx,
			`
			SELECT EXISTS(
				SELECT 1
				FROM schema_migrations
				WHERE version=$1
			)
			`,
			version,
		).Scan(&exists)


		if err != nil {
			return err
		}


		if exists {
			fmt.Println(
				"migration already applied:",
				version,
			)

			continue
		}



		sqlFile, err := os.ReadFile(file)

		if err != nil {
			return err
		}



		_, err = db.ExecContext(
			ctx,
			string(sqlFile),
		)


		if err != nil {
			return err
		}



		_, err = db.ExecContext(
			ctx,
			`
			INSERT INTO schema_migrations(version)
			VALUES($1)
			`,
			version,
		)


		if err != nil {
			return err
		}


		fmt.Println(
			"migration applied:",
			version,
		)
	}


	return nil
}