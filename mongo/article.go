package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//文章相关的
type Article struct {
	UUID     string
	Title    string
	Context  string
	UserUUID string
}

var article *mgo.Collection

func init() {
	initMongo()
	article = db.C("article")
}

//Add user
func (m *Article) Insert() (err error) {
	err = article.Insert(m)
	return
}

//find by UUID
func (m *Article) GetByUUID() error {
	return article.Find(bson.M{"UUID": m.UUID}).One(m)
}
