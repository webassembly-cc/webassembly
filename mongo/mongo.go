package mongo

import (
	"log"

	"sync"

	"gopkg.in/mgo.v2"
)

var (
	mongoSession *mgo.Session  //session
	db           *mgo.Database //database
	so           sync.Once
)

//init mogondb
func initMongo() {
	//初始化一次  修复一个bug  就是包初始化的顺序 确保只会初始化一次
	so.Do(func() {
		var err error
		mongoSession, err = mgo.Dial("127.0.0.1")
		if err != nil {
			panic(err)
		}
		log.Println("init mongo success")
		mongoSession.SetMode(mgo.Monotonic, true)
		db = mongoSession.DB("webassembly")
	})
}
