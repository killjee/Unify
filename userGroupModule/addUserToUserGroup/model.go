package addUserToUserGroup

import(
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

var db *sql.DB

func getDb() *sql.DB {
	connStr := "user=postgres dbname=unify password=123456"
	db, err := sql.Open("postgres", connStr)
	if err != nil {	
		fmt.Println("Jhandu balm")
	}

	return db
}

func validateGroupUuidInDb(groupUuid string) bool {
	if (db == nil) {
		db = getDb()
	}

	sqlQuery := `SELECT uuid FROM UserGroupMetaData WHERE uuid = $1`

	var dump string
	fmt.Println(groupUuid)
	err := db.QueryRow(sqlQuery, groupUuid).Scan(&dump)
	if err == sql.ErrNoRows {
		return false
	}
	fmt.Println("JhunJhuna")
	return true
}

func addUsersToDb(groupUuid string, userList []string) bool {
	if (db == nil) {
		db = getDb()
	}

	var queryString string
	queryString = "INSERT INTO UserGroupMembers (groupuuid, useruuid) VALUES "
	numberUser := len(userList)
	fmt.Println(userList)
	for i:=0; i < numberUser; i++ {
		fmt.Println(i)
		fmt.Println(userList[i])
		tempString  := "('" + groupUuid + "', '" + userList[i] + "')"
		if (i != numberUser-1) {
			tempString += ", "
		} else {
			tempString += ";"
		}
		queryString += tempString
	}

	fmt.Println(queryString)
	_, err := db.Exec(queryString)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("Success")
	return true
}