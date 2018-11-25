package main

import (
	"fmt"
	"net/http"

	"github.com/cpekyaman/udemy_go_web/s15_gomongo/controller"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)

	uc := controller.NewUserController()
	uc.Register(r)

	http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome\n")
}
