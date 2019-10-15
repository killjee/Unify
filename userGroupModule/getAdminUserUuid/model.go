package getAdminUserUuid

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

func getAdminUuidFromDb(uuid string) (string, int){
	var db *sql.DB = getDb()

	sqlQuery := `SELECT adminUserUuid FROM UserGroupMetaData WHERE uuid = $1`

	var adminUuid string
	err := db.QueryRow(sqlQuery, uuid).Scan(&adminUuid)
	if err == sql.ErrNoRows {
		fmt.Println(err)
		return "", 404
	} else if err != nil {
		fmt.Println(err)
		return "", 500
	}
	return adminUuid, -1
}