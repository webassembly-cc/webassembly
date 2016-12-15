package main

import (
	"github.com/gorilla/mux"
	"webassembly.cc/webassembly/controllers"
)

func initRouter() *mux.Router {
	router := mux.NewRouter()

	//user begin router
	user := new(controllers.UserCtrl)

	router.HandleFunc("/api/user/regist", user.Regist)
	router.HandleFunc("/api/user/login", user.Login)

	//article
	article := new(controllers.ArticleCtrl)
	router.HandleFunc("/api/article/list", article.List)
	router.HandleFunc("/api/article/insert", article.Insert)

	return router
}
