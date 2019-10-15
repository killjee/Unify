package getParentUserGroup

import(
	"encoding/json"
	"net/http"
	"fmt"
)

type Response struct {
	ParentUuid string `json:"parentUuid"`
}

func RequestHandler(response http.ResponseWriter, Request *http.Request) {
	var request map[string]interface{}
	json.NewDecoder(Request.Body).Decode(&request)

	var uuid string
	if val, ok := request["uuid"]; ok {
		uuid = fmt.Sprintf("%v", val)
	} else {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("400 - Bad request"))
		return
	}

	responseUuid, err := getParentUserGroup(uuid)
	fmt.Println(responseUuid)
	if err != -1 {
		if err == 404 {
			response.WriteHeader(http.StatusNotFound)
			response.Write([]byte("404 - Group not found"))
			return
		} else {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte("500 - Something bad happened"))
			return
		}
	} else {
		response.Header().Set("Content-Type", "application/json")
		resp := Response { ParentUuid: responseUuid }
		fmt.Println(resp)
		json.NewEncoder(response).Encode(resp)
		return
	}
}