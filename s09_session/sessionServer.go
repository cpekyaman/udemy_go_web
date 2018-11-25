package main

import (
	"fmt"
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/cpekyaman/udemy_go_web/s09_session/domain"
	"github.com/cpekyaman/udemy_go_web/s09_session/store"
)

var tpl *template.Template
var sessions store.SessionStore
var users store.UserStore

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	sessions.Init()
	users.Init()
}

func main() {
	http.HandleFunc("/", login)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/show", show)

	http.ListenAndServe(":8080", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	if sessions.IsLoggedIn(r) {
		fmt.Println("User is already logged in")
		http.Redirect(w, r, "/show", http.StatusSeeOther)
		return
	}

	sid := sessions.GetOrCreate(w, r)

	if r.Method == http.MethodPost {
		ok, eMsg := sessions.Login(sid, r, &users)
		if !ok {
			tpl.ExecuteTemplate(w, "login.gohtml", eMsg)
		} else {
			http.Redirect(w, r, "/show", http.StatusSeeOther)
		}
		return
	}

	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func signup(w http.ResponseWriter, r *http.Request) {
	if sessions.IsLoggedIn(r) {
		fmt.Println("User is already logged in")
		http.Redirect(w, r, "/show", http.StatusSeeOther)
		return
	}

	sid := sessions.GetOrCreate(w, r)

	if r.Method == http.MethodPost {
		uname := r.FormValue("userName")
		fname := r.FormValue("firstName")
		lname := r.FormValue("lastName")
		pwd := r.FormValue("password")

		u, ok := users.FindByUserName(uname)
		if ok {
			fmt.Println("UserName already exists")
			tpl.ExecuteTemplate(w, "signup.gohtml", map[string]string{
				"userName":  uname,
				"firstName": fname,
				"lastName":  lname,
				"err":       "UserName already exists",
			})
			return
		}

		hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("Could not hash password", err.Error())
			http.Redirect(w, r, "/signup", http.StatusInternalServerError)
			return
		}

		u = domain.User{
			UserName:  uname,
			Password:  hashedPwd,
			FirstName: fname,
			LastName:  lname,
		}
		sessions.SetUser(sid, uname)
		users.Save(u)

		fmt.Println("created user", u, "in session", sid)
		http.Redirect(w, r, "/show", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", map[string]string{})
}

func show(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(store.SIDCookieName)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	uname, ok := sessions.GetUser(c.Value)
	if !ok {
		fmt.Println("User not found in session", c.Value)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Println("found user", uname, "in session", c.Value)

	usr, ok := users.FindByUserName(uname)
	if !ok {
		fmt.Println("User not found in user store", uname)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "show.gohtml", usr)
}
