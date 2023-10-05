package session

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/url"
	"time"

	//"bytes"
	"encoding/gob"
	//"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	//"time"

	//"fincal/internal/cookies"
	"fincal/pkg/models"

	"github.com/gorilla/sessions"
)

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

type Session interface {
	Set(key, value interface{}) error //set session value
	Get(key interface{}) interface{}  //get session value
	Delete(key interface{}) error     //delete session value
	SessionID() string                //back current sessionID
}

// Store the cookie store which is going to store session data in the cookie
var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

var provides = make(map[string]Provider)

func init() {
	gob.Register(&models.User{})
}

type Manager struct {
	CookieName  string     //private cookiename
	Lock        sync.Mutex // protects session
	Provider    Provider
	MaxLifetime int64
}

func NewManager(provideName, cookieName string, maxLifetime int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{
		Provider:    provider,
		CookieName:  cookieName,
		MaxLifetime: maxLifetime,
	}, nil
}

// Register makes a session provider available by the provided name.
// If a Register is called twice with the same name or if the driver is nil,
// it panics.
func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provider " + name)
	}
	provides[name] = provider
}

func (manager *Manager) GC() {
	manager.Lock.Lock()
	defer manager.Lock.Unlock()
	manager.Provider.SessionGC(manager.maxlifetime)
	time.AfterFunc(time.Duration(manager.maxlifetime), func() { manager.GC() })
}

func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.Lock.Lock()
	defer manager.Lock.Unlock()
	cookie, err := r.Cookie(manager.CookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		session, _ = manager.Provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.CookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.MaxLifetime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.Provider.SessionRead(sid)
	}
	return
}

// Destroy sessionid
func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.CookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		manager.Lock.Lock()
		defer manager.Lock.Unlock()
		_ = manager.Provider.SessionDestroy(cookie.Value)
		expiration := time.Now()
		cookie := http.Cookie{Name: manager.CookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
		http.SetCookie(w, &cookie)
	}
}

// IsLoggedIn will check if the user has an active session and return True
func IsLoggedIn(r *http.Request) bool {
	sessionName := r.UserAgent() + ":" + r.RemoteAddr
	session, err := Store.Get(r, sessionName)
	if err != nil {
		log.Printf("IsLoggedIn: %s", err.Error())
		return false
	}

	if session.Values["loggedIn"] == "true" {
		////user, err := getUserCookie(r)
		////if err != nil {
		////	log.Printf("getUserCookie: %s", err.Error())
		////	return false
		////}
		//log.Printf("session user: %+v", user)
		//return true
	}
	return false
}

func SetLoggedIn(r *http.Request, w http.ResponseWriter, user *models.User) {
	sessionName := r.UserAgent() + ":" + r.RemoteAddr
	session, _ := Store.Get(r, sessionName)

	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 1,
		HttpOnly: true,
	}

	session.Values["loggedIn"] = "true"
	session.Values["user"] = &user
	err := session.Save(r, w)
	if err != nil {
		log.Printf("SetLoggedIn: %s", err.Error())
		return
	}

	//setUserCookie(w, user)
}

func SetLoggedOut(r *http.Request, w http.ResponseWriter) {
	session, _ := Store.Get(r, "session")
	session.Values["loggedIn"] = "false"
	session.Values["user"] = nil

	err := session.Save(r, w)
	if err != nil {
		log.Printf("SetLoggedOut: %s", err.Error())
		return
	}
}

//func getUserCookie(r *http.Request) (*models.User, error) {
//	value, err := cookies.ReadEncrypted(r, "user", []byte(os.Getenv("COOKIE_SECRET")))
//
//	var user = &models.User{}
//	var buf bytes.Buffer
//
//	buf.WriteString(value)
//	err = gob.NewDecoder(&buf).Decode(user)
//
//	if err != nil {
//		switch {
//		case errors.Is(err, http.ErrNoCookie):
//			return nil, fmt.Errorf("cookie not found")
//		case errors.Is(err, cookies.ErrInvalidValue):
//			return nil, fmt.Errorf("invalid cookie")
//		default:
//			log.Println(err)
//			return nil, fmt.Errorf("server error")
//		}
//	}
//	return user, nil
//}
//
//func setUserCookie(w http.ResponseWriter, user *models.User) {
//	// Initialize a buffer to hold the gob-encoded data.
//	var buf bytes.Buffer
//
//	// Gob-encode the user data, storing the encoded output in the buffer.
//	err := gob.NewEncoder(&buf).Encode(user)
//	if err != nil {
//		log.Printf("setUserCookie: %s", err.Error())
//		return
//	}
//
//	// Call buf.String() to get the gob-encoded value as a string and set it as the cookie value.
//	cookie := http.Cookie{
//		Name:     "user",
//		Value:    buf.String(),
//		Path:     "/",
//		MaxAge:   3600,
//		Expires:  time.Now().Add(time.Hour),
//		HttpOnly: true,
//		Secure:   true,
//		SameSite: http.SameSiteLaxMode,
//	}
//	err = cookies.WriteEncrypted(w, cookie, []byte(os.Getenv("COOKIE_SECRET")))
//	if err != nil {
//		log.Printf("setUserCookie: %s", err.Error())
//		return
//	}
//	_, err = w.Write([]byte("cookie set!"))
//	if err != nil {
//		log.Printf("setUserCookie: %s", err.Error())
//	}
//}
