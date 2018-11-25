package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type address struct {
	City string
	Town string
}

type user struct {
	UserName string `json:"uname"`
	Email    string
	Address  address `json:"addr"`
}

type person struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Items     []string
}

func main() {
	http.HandleFunc("/m", marshal)
	http.HandleFunc("/um", unmarshal)
	http.HandleFunc("/e", encode)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func unmarshal(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var u user
	json.Unmarshal(body, &u)
	fmt.Println(u)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status" : "OK"}`))
}

func marshal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p := person{
		"Alex",
		"Murphy",
		[]string{"Gun", "Armor"},
	}

	jsonData, err := json.Marshal(p)
	if err != nil {
		w.Write([]byte("{}"))
	} else {
		w.Write(jsonData)
	}
}

func encode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p := person{
		"Alex",
		"Murphy",
		[]string{"Gun", "Armor"},
	}

	json.NewEncoder(w).Encode(p)
}
