package postgresql

import (
	"cdnetwork/pkg/util"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func GetAgents() {
	config := util.GetConfig()
	dbaccess, err := NewDataAccessLayer2(config)

	if err != nil {
		panic(err)
	}
	// defer dbaccess.Listener.Close()
	// defer dbaccess.SshClient.Close()
	defer dbaccess.Db.Close()

	fmt.Println("start query")

	rows, err := dbaccess.Db.Query("SELECT id, username FROM dbo.agents")
	fmt.Println("start row")

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	fmt.Println("end query")

	for rows.Next() {
		var (
			id       uuid.UUID
			username string
		)
		// Scan the values from the row into variables
		err := rows.Scan(&id, &username)
		if err != nil {
			panic(err)
		}
		// Print the values
		fmt.Printf("ID: %d, Domain: %s\n", id, username)
	}
	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		panic(err)
	}
}
