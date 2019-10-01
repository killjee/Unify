package createUserGroup

import(
	"net/http"
	"fmt"
)

func RequestHandler(w http.ResponseWriter, Request *http.Request) {
	fmt.Print(w, "Hello World")
}