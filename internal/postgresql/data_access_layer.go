package postgresql

import (
	"cdnetwork/pkg/log"
	"cdnetwork/pkg/util"
	"database/sql"
	"fmt"
)

func NewDataAccessLayer(config util.TgsConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=prefer",
		config.Postgresql.PgUsername,
		config.Postgresql.PgPassword,
		config.Postgresql.PgHost,
		config.Postgresql.PgDb,
	)

	log.LogInfo("conn:" + connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.LogFatal(fmt.Sprintf("Failed to connect to database: %v", err))
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.LogFatal(fmt.Sprintf("Failed to ping database: %v", err))
		panic(err)
	}

	return db, nil
}
