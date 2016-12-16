package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"webassembly.cc/webassembly/mongo"
)

type ArticleCtrl struct {
}

//添加文章的 测试数据而已
func (c *ArticleCtrl) Insert(w http.ResponseWriter, r *http.Request) {
	m := new(mongo.Article)
	m.UUID = "11111"
	m.Title = "name"
	m.Context = "Context"

	err := m.Insert()

	if err != nil {
		fmt.Fprintln(w, err)
	} else {
		fmt.Fprintln(w, m.UUID)
	}

}

//list user
func (c *ArticleCtrl) List(w http.ResponseWriter, r *http.Request) {
	m, err := mongo.GetArticleList()

	if err != nil {
		fmt.Fprintln(w, "error:", err)
	} else {
		j, _ := json.Marshal(m)
		fmt.Fprintln(w, string(j))
	}
}

//update user
func (c *ArticleCtrl) Update(w http.ResponseWriter, r *http.Request) {
	m, err := mongo.ArticleFind()

	if err != nil {
		fmt.Fprintln(w, "error:", err)
	} else {
		j, _ := json.Marshal(m)
		fmt.Fprintf(w, "%s", j)
	}

}

func (c *ArticleCtrl) Delete(w http.ResponseWriter, r *http.Request) {

}
