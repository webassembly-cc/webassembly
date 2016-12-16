package mongo

import (
	"sync"

	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//文章相关的
type Article struct {
	UUID     string `json:"uuid"`
	Title    string `json:"title"`
	Context  string `json:"context"`
	UserUUID string `json:"useruuid"`
}

var article *mgo.Collection
var syOne sync.Once

func init() {
	syOne.Do(func() {
		initMongo()
		article = db.C("article")
		log.Println("article init")
	})

}

//Add user
func (m *Article) Insert() (err error) {

	err = article.Insert(m)
	return
}

//find by UUID
func (m *Article) GetByUUID() error {
	return article.Find(bson.M{"uuid": m.UUID}).One(m)
}

func ArticleFind() (Article, error) {
	m := Article{}
	err := article.Find(bson.M{"uuid": "11111"}).One(&m)
	if err != nil {
		return m, err
	}
	return m, nil
}

func GetArticleList() (list []Article, err error) {
	err = article.Find(&bson.M{}).All(&list)
	return
}
