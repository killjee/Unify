package createUserGroup

import(
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

func getDb() *sql.DB {
	connStr := "user=postgres dbname=unify password=123456"
	db, err := sql.Open("postgres", connStr)
	if err != nil {	
		fmt.Println("Jhandu balm")
	}

	return db
}

func addGroupToDb (uuid string, name string, adminUserUuid string, parentGroupUuid string ) bool {
	var db *sql.DB = getDb()

	sqlQuery := `INSERT INTO UserGroupMetaData (uuid, name, adminUserUuid, parentGroupUuid)
				 VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(sqlQuery, uuid, name, adminUserUuid, parentGroupUuid)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}