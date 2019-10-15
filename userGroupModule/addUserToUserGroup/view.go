package addUserToUserGroup

import(
	"encoding/json"
	"net/http"
	"fmt"
	"strings"
)

func validateRequest(groupUuid string, userList []string) bool {

	//ToDo(Satyam): Validate userList once userTable is made
	if (!validateGroupUuid(groupUuid)) {
		return false
	}

	return true
}

func RequestHandler(response http.ResponseWriter, Request *http.Request) {
	var request map[string]interface{}
	json.NewDecoder(Request.Body).Decode(&request)

	userList := fmt.Sprintf("%v", request["userUuid"])
	groupUuid := fmt.Sprintf("%v", request["groupUuid"])

	userUuid := strings.Split(userList, ",")

	if (!validateRequest(groupUuid,userUuid)) {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("400 - Bad request"))
		return
	}

	us := userList[len(userList) - 1]
	fmt.Println(us)

	isAdded := addUsers(groupUuid, userUuid)
	fmt.Println(isAdded)
	if (!isAdded) {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("500 - Something bad happened"))
	} else {
		response.WriteHeader(http.StatusOK)
		response.Write([]byte("200 - Added"))
	}
}