package getParentUserGroup

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

func getParentUserGroupFromDb(uuid string) (string, int){
	var db *sql.DB = getDb()

	sqlQuery := `SELECT parentGroupUuid FROM UserGroupMetaData WHERE uuid = $1`

	var parentUuid string
	err := db.QueryRow(sqlQuery, uuid).Scan(&parentUuid)
	if err == sql.ErrNoRows {
		fmt.Println(err)
		return "", 404
	} else if err != nil {
		fmt.Println(err)
		return "", 500
	}
	
	return parentUuid, -1
}