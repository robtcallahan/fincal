package sessions

import (
	"fincal/pkg/models"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

// Store the cookie store which is going to store session data in the cookie
var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

// IsLoggedIn will check if the user has an active session and return True
func IsLoggedIn(r *http.Request) bool {
	session, _ := Store.Get(r, "session")
	if session.Values["loggedin"] == "true" {
		return true
	}
	return false
}

func SetLoggedIn(r *http.Request, user *models.User) {
	session, _ := Store.Get(r, "session")
	session.Values["loggedin"] = "true"
	session.Values["loggedin"] = "true"
	session.Values["loggedin"] = "true"
}
