package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"webassembly.cc/webassembly/mongo"
)

type ArticleCtrl struct {
}

func (c *ArticleCtrl) Insert(w http.ResponseWriter, r *http.Request) {
	m := new(mongo.Article)
	m.UUID = "11111"
	m.Title = "name"
	m.Context = "Context"

	err := m.Insert()
	if err != nil {
		w.Write([]byte("err"))

	} else {
		w.Write([]byte(m.UUID))
	}
}

//list user
func (c *ArticleCtrl) List(w http.ResponseWriter, r *http.Request) {
	m := new(mongo.Article)
	m.UUID = "11111"
	err := m.GetByUUID()
	if err != nil {
		fmt.Fprintln(w, err)
	} else {
		j, err := json.Marshal(m)
		fmt.Fprintln(w, j, err)
	}
}

//update user
func (c *ArticleCtrl) Update(w http.ResponseWriter, r *http.Request) {

}

func (c *ArticleCtrl) Delete(w http.ResponseWriter, r *http.Request) {

}
