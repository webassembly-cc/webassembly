package controllers

import (
	"net/http"

	"fmt"

	"encoding/json"

	"webassembly.cc/webassembly/mongo"
)

type UserCtrl struct {
	Controller
}

//Login
func (c *UserCtrl) Login(w http.ResponseWriter, r *http.Request) {

}

//Regist
func (c *UserCtrl) Regist(w http.ResponseWriter, r *http.Request) {

	user := new(mongo.User)
	user.UUID = "11111"
	user.Name = "name"
	user.Password = "pasword"

	err := user.Insert()
	if err != nil {
		w.Write([]byte("err"))

	} else {
		w.Write([]byte(user.UUID))
	}

}

//list user
func (c *UserCtrl) List(w http.ResponseWriter, r *http.Request) {

}

//list user
func (c *UserCtrl) Info(w http.ResponseWriter, r *http.Request) {
	user := new(mongo.User)
	user.UUID = "11111"
	err := user.GetByUUID()
	if err != nil {
		fmt.Fprintln(w, err)
	} else {
		j, err := json.Marshal(user)
		fmt.Fprintln(w, j, err)
	}

}

//update user
func (c *UserCtrl) Update(w http.ResponseWriter, r *http.Request) {

}
