package createUserGroup

import(
	"encoding/json"
	"net/http"
	"fmt"
)

func checkAndGetValueOfKey(key string, request *map[string]interface{}) string {
	if value, ok := (*request)[key]; ok {
		return fmt.Sprintf("%v", value)
	} else {
		fmt.Println("lol")
		return "nil"
	}
}

func validateRequest(request *map[string]interface{}, validatedRequest *map[string]string) {

	(*validatedRequest)["uuid"] = checkAndGetValueOfKey("uuid", request)
	(*validatedRequest)["name"] = checkAndGetValueOfKey("name", request)
	(*validatedRequest)["adminUserUuid"] = checkAndGetValueOfKey("adminUserUuid", request)
	(*validatedRequest)["parentGroupUuid"] = checkAndGetValueOfKey("parentGroupUuid", request)

	//ToDo(Satyam): Add validation for valid adminUserUuid and ParentGroupUuid
}

func RequestHandler(response http.ResponseWriter, Request *http.Request) {
	var request map[string]interface{}
	json.NewDecoder(Request.Body).Decode(&request)

	var validatedRequest map[string]string
	validatedRequest = make(map[string]string)
	validateRequest(&request, &validatedRequest)

	if validatedRequest["uuid"] == "nil" || validatedRequest["name"] == "nil" || validatedRequest["adminUserUuid"] == "nil" || validatedRequest["parentGroupUuid"] == "nil" {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("400 - Bad request"))
		return
	}

	isSuccess := createGroup(validatedRequest["uuid"], validatedRequest["name"],
								 validatedRequest["adminUserUuid"],
								 validatedRequest["parentGroupUuid"])
	if !isSuccess {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("500 - Something bad happened"))
		return
	} else {
		response.WriteHeader(http.StatusCreated)
		response.Write([]byte("201 - User group created successfully"))
		return
	}
}