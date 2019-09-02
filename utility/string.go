package utility

import (
	"fmt"
	"net/http"

	"github.com/gorilla/context"
)

// GetLoggedInUserID returns current loggedIn user id
// it reads userId from request context and
// return id as string.
// The loggedIn user id was set in context while
// validating JWT token from request header
func GetLoggedInUserID(r *http.Request) string {
	id := context.Get(r, "userId")
	return fmt.Sprintf("%v", id)
}
