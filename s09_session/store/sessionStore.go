package store

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	SIDCookieName = "JSESSIONID"
)

type SessionStore struct {
	store map[string]string
}

func (s *SessionStore) Init() {
	s.store = make(map[string]string)
}

func (s *SessionStore) GetOrCreate(w http.ResponseWriter, r *http.Request) string {
	c, err := r.Cookie(SIDCookieName)

	if err != nil {
		sid, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  SIDCookieName,
			Value: sid.String(),
		}
		http.SetCookie(w, c)
		fmt.Println("Created session", c.Value)
	} else {
		fmt.Println("Using existing session", c.Value)
	}

	return c.Value
}

func (s *SessionStore) SetUser(sid string, uname string) {
	s.store[sid] = uname
}

func (s *SessionStore) GetUser(sid string) (string, bool) {
	uname, ok := s.store[sid]
	return uname, ok
}

func (s *SessionStore) Login(sid string, r *http.Request, users *UserStore) (bool, string) {
	uname := r.FormValue("userName")
	pwd := r.FormValue("password")

	user, ok := users.FindByUserName(uname)
	if !ok {
		fmt.Println("User", uname, "not found")
		return false, "User Not Found"
	}

	err := bcrypt.CompareHashAndPassword(user.Password, []byte(pwd))
	if err != nil {
		fmt.Println("Login failed for user", uname)
		return false, "Login Failed"
	}

	s.SetUser(sid, uname)
	return true, ""
}

func (s *SessionStore) IsLoggedIn(r *http.Request) bool {
	c, err := r.Cookie(SIDCookieName)
	if err != nil {
		return false
	}

	_, ok := s.GetUser(c.Value)
	return ok
}
