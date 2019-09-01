package utility

import (
	"fmt"
	"net/http"

	"github.com/gorilla/context"
)

func GetLoggedInUserId(r *http.Request) string {
	id := context.Get(r, "userId")
	return fmt.Sprintf("%v", id)
}
