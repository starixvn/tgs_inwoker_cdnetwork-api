package postgresql

import (
	"cdnetwork/pkg/log"
	"cdnetwork/pkg/util"
	"database/sql"
	"fmt"
)

func NewDataAccessLayer2(config util.TgsConfig) (*DataAccess, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=prefer",
		config.Postgresql.PgUsername,
		config.Postgresql.PgPassword,
		config.Postgresql.PgHost,
		config.Postgresql.PgDb,
	)

	fmt.Println("conn:" + connStr)

	db, err := sql.Open("postgres1", connStr)
	if err != nil {
		log.LogFatal(fmt.Sprintf("Failed to connect to database: %v", err))
		panic(err)
	}

	// fmt.Println("ping")

	// err = db.Ping()
	// if err != nil {
	// 	log.LogFatal(fmt.Sprintf("Failed to ping database: %v", err))
	// }

	// fmt.Println("pinged")

	return &DataAccess{
		Db: db,
	}, nil
}
