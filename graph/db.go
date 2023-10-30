// db.go

package graph

import (
	"database/sql"
	"fmt"
	"github.com/go-pg/pg"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func Connect() *pg.DB {
	connStr := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select version()")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var version string
	for rows.Next() {
		err := rows.Scan(&version)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("version=%s\n", version)

	opt, err := pg.ParseURL(connStr)
	if err != nil {
		panic(err)
	}

	//connStr := "postgres://hasura_role_6b3fbecd-176f-4b0a-9b93-34c9c8c73116:THLmKB6FS1YZ@proud-wave-70962163.us-east-2.aws.neon.tech/integral-ant-78_db_3188404"
	//opt, err := pg.ParseURL(connStr)
	//if err != nil {
	//	panic(err)
	//}
	dbs := pg.Connect(opt)
	if _, DBStatus := dbs.Exec("SELECT 1"); DBStatus != nil {
		fmt.Println(DBStatus)
	}
	return dbs
}
