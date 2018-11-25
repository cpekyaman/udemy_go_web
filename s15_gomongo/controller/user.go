package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"

	"github.com/cpekyaman/udemy_go_web/s15_gomongo/model"
	"github.com/julienschmidt/httprouter"
)

const (
	dbname = "udemy_go_web"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("view/*"))
}

type UserController struct {
	session *mgo.Session
}

func NewUserController() *UserController {
	s := connect()
	uc := &UserController{
		session: s,
	}

	return uc
}

func connect() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return s
}

func (uc UserController) Register(r *httprouter.Router) {
	r.POST("/users", uc.create)
	r.GET("/users/:id", uc.get)
	r.DELETE("/users/:id", uc.delete)
}

func (uc UserController) create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := model.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	uc.session.DB(dbname).C("users").Insert(u)

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	oid := bson.ObjectIdHex(id)

	u := model.User{}

	if err := uc.session.DB(dbname).C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB(dbname).C("users").RemoveId(oid); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
